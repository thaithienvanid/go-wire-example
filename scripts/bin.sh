#!/usr/bin/env sh

if [[ "${DEBUG}" =~ ^1|yes|true ]]; then
  echo "DEBUG=true"
  set -o xtrace
fi

SCRIPTPATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"

CURRENT_DIR=$SCRIPTPATH
ROOT_DIR="$(dirname $CURRENT_DIR)"
PROJECT_NAME="$(basename $ROOT_DIR)"

BUILD_DIR=${ROOT_DIR}/_build
VERSION=$(git rev-parse --short HEAD)

function build {
  echo "BUILD ..."
  echo "VERSION ${VERSION}"

  rm -r ${BUILD_DIR} >/dev/null 2>&1
  mkdir _build

  go build -o ${BUILD_DIR}/main .

  echo "VERSION=${VERSION}" > ${BUILD_DIR}/.base.env

  # cp .env ${BUILD_DIR}/
  cp config.yaml ${BUILD_DIR}/
  cp ${SCRIPTPATH}/run.sh ${BUILD_DIR}/

  cd $RUNNING_DIR

  echo "BUILD DONE!"
}

function envup {
  echo "ENVUP ..."

  set -o allexport
    source ${ROOT_DIR}/.env
  set +o allexport

  export VERSION=${VERSION}

  echo "ENVUP DONE!"
}

function start() {
  echo "START ..."
  echo "VERSION ${VERSION}"

  envup
  go run . --config=${ROOT_DIR}/config.yaml

  echo "START DONE!"
}

function code_lint() {
  echo "LINT ..."

  command -v golangci-lint >/dev/null 2>&1 || {
    echo ""
    echo "project is installing golangci-lint"
    curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0
  }

  golangci-lint run --timeout 10m ./internal/... ./pkg/...

  echo "LINT DONE!"
}

function code_format() {
  echo "FORMAT ..."

  go generate ./...

  gofmt -w internal/ pkg/

  goimports -local go-wire-example -w internal/ pkg/

  echo "FORMAT DONE!"
}

function main() {
  case $1 in
    build)
      build ${@:2}
      ;;
    start)
      start ${@:2}
      ;;
    code_lint)
      code_lint ${@:2}
      ;;
    code_format)
      code_format ${@:2}
      ;;
    *)
      echo "build|start|code_lint|code_format"
      ;;
  esac
}

if [ "${1}" != "--source-only" ]; then
  main "${@}"
fi
