#!/bin/sh
WORK_DIR=$(pwd)
SCRIPT_DIR=$(dirname $0)
cd ${SCRIPT_DIR}
pwd
../node_modules/.bin/openapi-generator-cli generate -c config.yaml
cd ${WORK_DIR}