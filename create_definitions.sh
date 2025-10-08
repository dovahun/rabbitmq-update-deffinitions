#!/usr/bin/env bash

set -euxo pipefail

PATH_TO_VALUES=''
PATH_TO_RABBITMQ_DEFINITIONS_DIR=''
PATH_TO_MAKE_DEFINITION_FILE=''

render_manifest_file="manifest-rabbitmq.yaml"
definitions_file_in_base64="load_definition-draft-1.json"

help() {
  printf "
  -v: Путь до values.yaml для рендера манифеста.
  -d: Путь до директории с definitions.
  -f: Указать путь где будет создан файл с definition.
  -c: Провалидировать файл с definition, включение через указание флага.
  "
}

while getopts 'v:d:f:c' flag; do
  case "${flag}" in
    v) PATH_TO_VALUES="${OPTARG}" ;;
    d) PATH_TO_RABBITMQ_DEFINITIONS_DIR="${OPTARG}" ;;
    f) PATH_TO_MAKE_DEFINITION_FILE="${OPTARG}" ;;
    c) ;;
    *) help
       exit 1 ;;
  esac
done

renderHelmChart() {
  helm template \
    -f $PATH_TO_VALUES \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/bindings.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/exchanges.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/global_parameters.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/parameters.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/permissions.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/policies.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/queues.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/users.yaml \
    -f $PATH_TO_RABBITMQ_DEFINITIONS_DIR/vhosts.yaml $PATH_TO_RABBITMQ_DEFINITIONS_DIR/../ > $render_manifest_file
}

makeDefinitionFile() {
  yq -r '.data["definitions.json"] // empty'  $render_manifest_file > $definitions_file_in_base64

  cat $definitions_file_in_base64 | base64 -d > $PATH_TO_MAKE_DEFINITION_FILE

  rm $render_manifest_file $definitions_file_in_base64
}

lintDefinitionFile(){
  if cat $PATH_TO_MAKE_DEFINITION_FILE | jq . ; then
    echo "PARSED JSON SUCCESSFULLY!"
    exit 0
  else
    echo "FAILED TO PARSE JSON, OR GOT FALSE/NULL"
    exit 1
  fi
}

renderHelmChart
makeDefinitionFile
if [[ $* == *-c* ]]; then
  lintDefinitionFile
fi
