#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 

# build on the native or default platform
echo "building native platform"


# generate kallax model first

echo "Generating model"
cd model
./build_and_generate_model.sh
cd ..

echo "Generating main"
if ! go generate ; then
   go build
   exit 1
fi

echo "Test main"
go test
echo "Build main"
go build
./kallax_example