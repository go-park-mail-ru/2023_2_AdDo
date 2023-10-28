#!/bin/bash

if [ ! -d "/home/$USER/db-data" ]; then
    echo "Database dir does not exist"
    exit 0
fi

echo "Database cleaning..."
sudo rm -r /home/$USER/db-data
if [ $? -ne 0 ]; then
    echo "Error database cleaning"
    exit 1
fi

echo "Database clean successful"
