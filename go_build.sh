#!/bin/sh -e

if [ ! -e third_party/src/github.com/rrramiro/azurectl ] ; then
	ln -s ../../../.. third_party/src/github.com/rrramiro/azurectl
fi

go run third_party.go install github.com/rrramiro/azurectl
