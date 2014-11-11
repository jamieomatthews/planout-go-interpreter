#!/bin/bash

LOCATION=test/*.txt

for f in $LOCATION
do
	echo ${f##*/}
	./interpreter ${f##*/}
done
