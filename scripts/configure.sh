#!/usr/bin/sh
# configure script just copies a given file list to test/
# directory that will be build into Docker container
[ -n "$@" ] && cp -vf "$@" test/ || echo "Any files to copy"