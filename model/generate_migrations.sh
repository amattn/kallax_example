#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 

kallax migrate --input . --out ./migrations --name initial_schema