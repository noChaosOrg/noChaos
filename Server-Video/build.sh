#!/bin/bash

sh clear.sh

export GO111MODULE=on

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/api
go build -o ../bin/api

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/scheduler
 go build -o ../bin/scheduler

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/streamserver
go build -o ../bin/streamserver

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/web
go build -o ../bin/web

echo "build finished"
