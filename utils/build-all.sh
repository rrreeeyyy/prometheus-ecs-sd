#!/usr/bin/env bash
set -e

PROJECT_ROOT=$(git rev-parse --show-toplevel)
VERSION=$(git describe --tags --dirty)
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")
BUILD_PLATFORM=$(uname -a | awk '{print tolower($1);}')
IMPORT_DURING_SOLVE=${IMPORT_DURING_SOLVE:-false}

if [[ "$(pwd)" != "${PROJECT_ROOT}" ]]; then
  echo "you are not in the root of the repo" 1>&2
  echo "please cd to ${PROJECT_ROOT} before running this script" 1>&2
  exit 1
fi

GO_BUILD_CMD="go build -a -installsuffix cgo"
GO_BUILD_LDFLAGS="-s -w -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${DATE} -X main.version=${VERSION} -X main.flagImportDuringSolve=${IMPORT_DURING_SOLVE}"

if [[ -z "${BUILD_PLATFORMS}" ]]; then
    BUILD_PLATFORMS="linux darwin"
fi

if [[ -z "${BUILD_ARCHS}" ]]; then
    BUILD_ARCHS="amd64 386"
fi

mkdir -p "${PROJECT_ROOT}/release"

for OS in ${BUILD_PLATFORMS[@]}; do
  for ARCH in ${BUILD_ARCHS[@]}; do
    NAME="prometheus-ecs-hako-sd-${OS}-${ARCH}"
    if [[ "${OS}" == "windows" ]]; then
      NAME="${NAME}.exe"
    fi

    if [[ "${OS}" == "darwin" && "${BUILD_PLATFORM}" == "darwin" ]]; then
      CGO_ENABLED=1
    else
      CGO_ENABLED=0
    fi

    echo "Building for ${OS}/${ARCH} with CGO_ENABLED=${CGO_ENABLED}"
    GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=${CGO_ENABLED} ${GO_BUILD_CMD} -ldflags "${GO_BUILD_LDFLAGS}"\
     -o "${PROJECT_ROOT}/release/${NAME}"
    shasum -a 256 "release/${NAME}" > "${PROJECT_ROOT}/release/${NAME}".sha256
  done
done
