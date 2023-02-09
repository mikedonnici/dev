#!/bin/bash
set -o errexit
set -o xtrace
set -o nounset
pwd
cd $HOME
cd -
echo $DOESNOTEXIST
echo "should not get here"
