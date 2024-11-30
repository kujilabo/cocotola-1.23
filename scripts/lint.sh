#!/bin/bash

# echo cocotola-web
# pushd cocotola-web
# npm run lint
# popd
# code=$?
# if [ $code -ne 0 ]; then
#   exit 1
# fi

task lint
code=$?
if [ $code -ne 0 ]; then
  exit 1
fi
