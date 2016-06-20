#!/bin/bash

APP="summon-conjur"
WORKDIR="/go/src/github.com/conjurinc/${APP}"

docker run --rm \
-v "$PWD":$WORKDIR \
-w $WORKDIR \
golang:1.6 \
bash -ceo pipefail "go get -u github.com/jstemmer/go-junit-report && \
go test -v $(go list ./... | grep -v /vendor/) | tee test.tmp \
&& cat test.tmp | go-junit-report > junit.xml && rm test.tmp"
