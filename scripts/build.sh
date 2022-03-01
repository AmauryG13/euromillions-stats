#!/bin/sh
set -e

MODULE="$1"

if [ -z $2 ]; then
    BUILD_ROOT_DIR=/${MODULE}
fi

GO_BUILD_EXEC=".${BUILD_ROOT_DIR}/ems-${MODULE}"
rm ${GO_BUILD_EXEC}

GO_BUILD_PREFIX="CGO_ENABLED=0"
GO_BUILD_FLAGS="-a --installsuffix cgo"

GO_BUILD_OUTPUT="-o ${GO_BUILD_EXEC}"
GO_BUILD_PACKAGES=".${BUILD_ROOT_DIR}/cmd/${MODULE}/main.go"


export ${GO_BUILD_PREFIX}
go build ${GO_BUILD_FLAGS} ${GO_BUILD_OUTPUT} ${GO_BUILD_PACKAGES}