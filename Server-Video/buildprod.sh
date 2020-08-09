#!/bin/bash

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/publish/api

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/publish/scheduler

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/publish/streamserver

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/publish/web

echo "buildprod finished"
