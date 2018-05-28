#! /bin/bash

getent passwd golang  > /dev/null

if [ $? -ne 0 ]; then
    echo "Creating user golang."
    adduser -r -s /sbin/nologin ${rpm.username}
fi

mkdir -p /var/log/${project.artifactId}