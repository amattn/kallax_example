#!/bin/sh

# the .gitignore file includes the line export_env_vars*
# so in order to prevent checking in secrets into git, please 
# make a copy of this file and name it export_env_vars.sh 
# or export_env_vars_foo.sh, etc.
# Then you can feel free to edit and add your actual passwords, hosts, etc.

# in bash, in order to export these variables you have to use the dot command:
#
# . export_env_vars.sh

export DBUSER="dbuser"
export DBPASS="haha you wish"
#export DBHOST=""
#export DBPORT=""
#export DBNAME=""