#!/bin/bash

# Get all generated schema files from ent
entSchemaDir=internal/ent/schema
graphSchemaDir=schema
schemas=$(find $entSchemaDir -name '*.go')

for file in $schemas
do
    file=${file##*/}
    schema=${file%.*}
    if [ -f "$graphSchemaDir/$schema.graphql" ]
    then
        # Check if file already exists
        if [ -f "$graphSchemaDir/$schema.graphql" ]
        then
            echo "$graphSchemaDir/$schema.graphql already exists, not regenerating."
        else

            touch $graphSchemaDir/$schema.graphql

            export name="${schema}"

            # Object is capitilized
            first=`echo $name|cut -c1|tr [a-z] [A-Z]`
            second=`echo $name|cut -c2-`
            export object="${first}${second}"

            # Generate a base graphql schema
            gomplate -f scripts/templates/graph.tpl > $graphSchemaDir/$schema.graphql

            echo +++ file created $graphSchemaDir/$schema.graphql
        fi
    fi
done