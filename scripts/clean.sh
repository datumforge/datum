#!/bin/bash
set -ueo pipefail
# Find all occurances of `go-template` and its variations and replace with current directory name
appName=go-template
appVariation=${appName/"-"/""}

#what we want to update to
newAppName=${PWD##*/}
newAppVariation=${newAppName/"-"/""}

echo +++ Update repo occurances with $newAppName
git grep -lz $appName | xargs -0 sed -i '' -e "s/${appName}/${newAppName}/g"

echo +++ Update variations with $appVariation
git grep -lz $appVariation | xargs -0 sed -i '' -e "s/${appVariation}/${newAppVariation}/g"