#!/bin/bash

echo "Preparing environments.."
docker-compose -f deployments/test/docker-compose.yml up -d
echo "Preparing successful done"

echo "Running tests with empty database..."
python3 test/testsuite/sign_up_test.py
python3 test/testsuite/music_test.py
python3 test/testsuite/me_test.py
python3 test/testsuite/logout_test.py
python3 test/testsuite/login_test.py
python3 test/testsuite/listen_test.py
python3 test/testsuite/auth_test.py
if [ $? -ne 0 ]; then
    echo "Error integration tests"
    exit 1
fi

echo "Tests passed"

docker-compose -f deployments/test/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error delete test environment"
    exit 1
fi
echo "Test enviroment deleted"
