#!/bin/bash
# chkconfig: 345 99 01
# description: ${project.artifactId} service
# Source function library.
. /etc/init.d/functions

start() {
    if [ -e /var/lock/subsys/${project.artifactId} ]; then
        echo -e '\e[31m [ERROR] \e[0m -> Service already running'
        exit 1
    fi

    echo -e 'Starting ${project.artifactId}:  [ \e[92m OK \e[0m ]'
    sh /opt/${project.artifactId}/bin/run.sh
    touch /var/lock/subsys/${project.artifactId}
}

stop() {
    if [ ! -e /var/lock/subsys/${project.artifactId} ]; then
        echo -e '\e[31m [ERROR] \e[0m -> Service is not running'
    else
        echo -e 'Shutting down ${project.artifactId}:  [ \e[92m OK \e[0m ]'
        kill `cat /var/run/${project.artifactId}.pid`
        rm -f /var/lock/subsys/${project.artifactId}
    fi
}

status() {
    if [ -e /var/lock/subsys/${project.artifactId} ]; then
        isRunning=$(ps --pid `cat /var/run/${project.artifactId}.pid` | wc -l)
        if [ "$isRunning" == "2" ]; then
            echo ${project.artifactId} is running, pid=`cat /var/run/${project.artifactId}.pid`
            exit 0
        fi
    fi

    echo ${project.artifactId} is not running
    exit 1
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    status)
        status
        ;;
    restart)
        stop
        sleep 2
        start
        ;;
    *)
        echo "Usage: ${project.artifactId} {start|stop|status|restart}"
        exit 1
        ;;
esac
exit $?