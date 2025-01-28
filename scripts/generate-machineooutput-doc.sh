#!/bin/bash

set -e

asciioutput() {
  OUTPUT=`${@}`
  ALLOUTPUT="== ${@}

[source,json]
----
${OUTPUT}
----
"
  echo ${ALLOUTPUT} > test.adoc
}

tmpdir=`mktemp -d`
cd $tmpdir
git clone https://github.com/openshift/nodejs-ex
cd nodejs-ex

# Commands that don't have json support
# app delete
# catalog describe
# catalog search
# component create
# component delete
# component link
# component log
# component push
# component unlink
# component update
# component watch
# config set
# config unset
# config view
# debug port-forward 
# preference set
# preference unset
# preference view
# service create
# service delete
# storage delete
# url create
# url delete
# login
# logout
# utils *
# version


# Alphabetical order for json output...

# Preliminary?
particle engine project delete foobar -f || true
sleep 5
particle engine project create foobar
sleep 5
particle engine create nodejs
particle engine push

# app
asciioutput particle engine app describe app -o json
particle engine app list -o json

# catalog
particle engine catalog list components -o json
particle engine catalog list services -o json

# component
particle engine component delete -o json
particle engine component push 

# project
particle engine project create foobar -o json
particle engine project delete foobar -o json
particle engine project list -o json

# service

## preliminary
particle engine service create mongodb-persistent mongodb --plan default --wait -p DATABASE_SERVICE_NAME=mongodb -p MEMORY_LIMIT=512Mi -p MONGODB_DATABASE=sampledb -p VOLUME_CAPACITY=1Gi
particle engine service list -o json

# storage
particle engine storage create mystorage --path=/opt/app-root/src/storage/ --size=1Gi -o json
particle engine storage list -o json
particle engine storage delete

# url
particle engine url create myurl
particle engine url list -o json
particle engine url delete myurl
