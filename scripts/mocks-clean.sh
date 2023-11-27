#!/bin/bash

echo "Mocks deleting..."

if [ ! -d "test/mocks" ]; then
    echo "Mocks does not exist"
    exit 0
fi

rm -r test/mocks
if [ $? -ne 0 ]; then
    echo "Error remove dir /test/mocks"
    exit 1
fi

echo "Mocks deleted successful"
