#!/usr/bin/env sh

if [[ "${DEBUG}" =~ ^1|yes|true ]]; then
  echo "DEBUG=true"
  set -o xtrace
fi

RUNNING_DIR="$(pwd -P)"

function envup {
  echo "ENVUP ..."

  set -o allexport
    source ${RUNNING_DIR}/.base.env
  set +o allexport

  echo "ENVUP DONE!"
}

function run {
  echo "RUN ..."

  chmod +x ${RUNNING_DIR}/main

  ${RUNNING_DIR}/main --config=${RUNNING_DIR}/config.yaml

  echo "RUN DONE!"
}

function main() {
  envup
  run
}

if [ "${1}" != "--source-only" ]; then
  main "${@}"
fi
