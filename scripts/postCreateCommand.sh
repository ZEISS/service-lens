#!/bin/bash
# This script is executed after the creation of a new project.

go install github.com/oapi-codegen/oapi-codegen/v2@latest
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/air-verse/air@latest
go install github.com/go-delve/delve/cmd/dlv@latest
