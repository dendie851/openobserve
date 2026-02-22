#!/bin/sh

# Ensure log directory exists
mkdir -p /var/log/app

# Start the application and redirect stdout/stderr to files
# We also use 'tee' if we want to see it in docker logs, 
# but the requirement is "pendekatan seperti di vm" (writing to files)
./main > /var/log/app/stdout.log 2> /var/log/app/stderr.log
