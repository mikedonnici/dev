#!/bin/bash
 # flags are: -a, -b 'option', -c
while getopts "ab:c:" opt
do
  case "$opt" in
  a) echo '-a flag';;
  b) echo "-b with arg ${OPTARG}";;
  c) echo "-c with arg ${OPTARG}";;
  esac
done
