#!/bin/bash
set -ueo pipefail
# Get all generated schema files from ent
repoRoot=$(git rev-parse --show-toplevel)
entSchemaDir=$repoRoot/internal/ent/schema
graphSchemaDir=schema
schemas=$(find $entSchemaDir -name '*.go')

# Track files to ignore the creation of graph schemas
skip_schemas=( $( cat ./scripts/files_to_skip.txt |tr "\n" " ") )

for file in $schemas
do
    file=${file##*/}
    schema=${file%.*}
    # Check if file already exists
    if [ -f "$graphSchemaDir/$schema.graphql" ]
    then
        echo "$graphSchemaDir/$schema.graphql already exists, not regenerating."
    elif [[ ! " ${skip_schemas[*]} " =~ " ${file} " ]]; then

        touch $graphSchemaDir/$schema.graphql

        echo creating $graphSchemaDir/$schema.graphql

        export name="${schema}"

        # Object is capitilized
        first=`echo $name|cut -c1|tr [a-z] [A-Z]`
        second=`echo $name|cut -c2-`
        export object="${first}${second}"

        # Generate a base graphql schema
        gomplate -f scripts/templates/graph.tpl > $graphSchemaDir/$schema.graphql

        echo +++ file created $graphSchemaDir/$schema.graphql
    fi
done
