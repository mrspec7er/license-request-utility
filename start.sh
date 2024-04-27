#!/bin/bash

go run script/database/main.go && go run script/migration/main.go

tail -f /dev/null