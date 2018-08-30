#!/usr/bin/env bash
cd /opt/${project.artifactId}
su -s /bin/sh ${rpm.username} -c "/opt/${project.artifactId}/${project.artifactId}" >>/var/log/${project.artifactId}/serfver.log 2>>/var/log/${project.artifactId}/server.log &
echo $!>/var/run/${project.artifactId}.pid