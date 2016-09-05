#!/bin/sh -e

CGO_ENABLED=0 go build --ldflags '-extldflags "-static"' -o bin/azurectl-linux64-static
