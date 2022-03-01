#!/bin/sh
set -e

MODULE="$1"

if [ -z $2 ]; then
    DOCKER_COMPOSE_FILE_DIR="."
else
    DOCKER_COMPOSE_FILE_DIR=".."
fi

DOCKER_COMPOSE_FILE_PATH="./${DOCKER_COMPOSE_FILE_DIR}/deployments/${MODULE}.yaml"

DOCKER_COMPOSE_CMD="up"
DOCKER_COMPOSE_ARGS="-d"

docker-compose -f ${DOCKER_COMPOSE_FILE_PATH} ${DOCKER_COMPOSE_CMD} ${DOCKER_COMPOSE_ARGS}