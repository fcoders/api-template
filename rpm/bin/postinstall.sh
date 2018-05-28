#! /bin/bash

chkconfig ${project.artifactId} on
echo "Changing owner of directories."
chown -R ${rpm.username}:${rpm.groupname} /opt/${project.artifactId}