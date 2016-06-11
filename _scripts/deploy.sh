#!/usr/bin/env bash
#
# Build and push Docker images to Docker Hub and quay.io.
#

cd "$(dirname "$0")" || exit 1

export IMAGE_PREFIX=deisci
docker login -e="$DOCKER_EMAIL" -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"
DEIS_REGISTRY='' make -C .. docker-build docker-push
# Here is a bogus change don't you know.
docker login -e="$QUAY_EMAIL" -u="$QUAY_USERNAME" -p="$QUAY_PASSWORD" quay.io
DEIS_REGISTRY=quay.io/ make -C .. docker-build docker-push
