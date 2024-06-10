#!/bin/bash
# This script is executed after the creation of a new project.

go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
go install github.com/air-verse/air@latest
go install github.com/go-delve/delve/cmd/dlv@latest
