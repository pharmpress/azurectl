#!/bin/sh -e

CGO_ENABLED=0 go build -a -tags netgo --ldflags '-w -extldflags "-static"' -o bin/azurectl-linux64-static
