#!/bin/bash

echo "Preparing environment.."
scripts/docker-services-strat.sh
if [ $? -ne "0" ]; then
    echo "Error prepare environment"
    exit 1
fi
echo "Preparing successful done"

cd test/testsuite && python3 run_test.py
if [ $? -ne 0 ]; then
    echo "Integration test failed"
    exit 1
fi
echo "Tests passed"

rm -r __pycache__ && cd ../.. && scripts/docker-services-down.sh
echo "Test enviroment deleted"
