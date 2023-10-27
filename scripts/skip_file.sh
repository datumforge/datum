#!/bin/bash
found=$(grep "^$1$" files_to_skip.txt)
if [ -z "$found" ]; then
    echo $1
fi