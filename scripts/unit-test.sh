#!/bin/bash

scripts/mocks-gen.sh
if [ $? -ne "0" ];then
    exit 1
fi

echo "Running unit tests..."
go test -coverprofile=all_files -coverpkg=./... ./...
if [ $? -ne "0" ];then
    echo "Unit tests failed"
    exit 1
fi
cat all_files | grep -v "cmd" | grep -v "test" | grep -v "init" | grep -v "/.*proto.*" | grep -v "mock" > testing_files
go tool cover -func=testing_files
rm testing_files all_files
rm -r test/mocks
