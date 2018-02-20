#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 

# generate kallax model first

echo "Generating model"
if ! rm kallax.go ; then
	# do nothing
	:
fi

if ! go generate ; then
	go build
    exit 1
fi

mkdir -p $GOPATH/pkg/darwin_amd64/github.com/amattn/
go build -o $GOPATH/pkg/darwin_amd64/github.com/amattn/kallax_example_model.a

./generate_migrations.sh

