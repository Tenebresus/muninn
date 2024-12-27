#!/bin/bash

muninnDir="/Users/tyler/Documents/programming/golang/projects/muninn"

/bin/bash $muninnDir/scripts/stop.sh

(docker run -it -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management) &
sleep 10
($muninnDir/bin/muninn_apiserver) &
($muninnDir/bin/muninn_scavenger) &
($muninnDir/bin/muninn_agent) &
($muninnDir/bin/muninn_incubator) &
