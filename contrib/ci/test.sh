#!/usr/bin/env bash

set -eoxf pipefail

JOB=$(docker run -d $1)
# let the registry run for a few seconds
sleep 5
# check that the registry is still up
docker ps -q --no-trunc=true | grep $JOB
docker rm -f $JOB
