#!/bin/bash

ls /Users/tyler/.docker/run/docker.sock

if [[ $? == 1 ]]; then
    echo "Please run docker desktop"
    exit 1
fi

docker stop rabbitmq

killall muninn_apiserver
killall muninn_agent
killall muninn_incubator
killall muninn_scavenger
