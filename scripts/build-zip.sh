#!/bin/bash

if [[ ! -e "build" ]]; then
    mkdir "build"
else
    rm -rf "build"
    mkdir "build"
fi

cp *.go ./build
cp go.mod ./build
cp go.sum ./build

# rm build.zip || echo '';
# (
#     cd build;
#     zip -r9 ../build.zip .
# )

