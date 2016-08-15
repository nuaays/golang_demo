#!/bin/bash

for file in `ls`
do
				newstr=`tr '[A-Z]' '[a-z]' <<< "$file"`
				mv $file $newstr
done

