#!/bin/bash

go fmt main.go interpreter.go operators.go utils.go
go build -o interpreter main.go interpreter.go operators.go utils.go
