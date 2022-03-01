#!/bin/sh
set -xe

MODULE="$1"

if [ -z $2 ]; then
    DOCKER_BUILD_DIR="."
else
    DOCKER_BUILD_DIR=".."
fi

DOCKER_FILE_DIR="./${DOCKER_BUILD_DIR}/build/docker/Dockerfile"

DOCKER_BUILD_NAME="amauryg13/ems-${MODULE}"
DOCKER_BUILD_TAG="latest"

DOCKER_BUILD_ARGS="--build-arg MODULE=${MODULE}"

docker build -f ${DOCKER_FILE_DIR} -t ${DOCKER_BUILD_NAME}:${DOCKER_BUILD_TAG} ${DOCKER_BUILD_ARGS} ${DOCKER_BUILD_DIR}

