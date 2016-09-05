#!/bin/sh -e

if [ ! -e third_party/src/github.com/pharmpress/azurectl ] ; then
	ln -s ../../../.. third_party/src/github.com/pharmpress/azurectl
fi

go run third_party.go install github.com/pharmpress/azurectl
