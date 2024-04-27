#!/bin/bash

sleep 2

go run script/database/main.go

sleep 3

go run script/migration/main.go

tail -f /dev/null