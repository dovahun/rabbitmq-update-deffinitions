#!/usr/bin/env sh

TMP_FILE=$(mktemp "${TMPDIR:-/tmp/}temp.XXXXXX")
CI_PATH=".."

echo ${TMP_FILE}

for entry in "$CI_PATH"/*.yaml
do
  KIND=$( basename ${entry} | sed 's/.yaml$//g')
  CHECK_FILE="${CI_PATH}/${KIND}.yaml"

  echo "Check file: $CHECK_FILE"

  echo "[" > ${TMP_FILE} && \
  cat ${CHECK_FILE} | yq r - definitions.${KIND}  >> ${TMP_FILE} && \
  echo "]" >> ${TMP_FILE} && \
  cat ${TMP_FILE} | jq

  if [ $? -ne 0 ]
  then
    rm -f ${TMP_FILE}
    echo "Checked $CHECK_FILE FAILED!"
    exit 1
  fi
   echo "Checked $CHECK_FILE SUCCESS!"
   rm -f ${TMP_FILE}
done

echo "ALL FILES DONE !!!"

