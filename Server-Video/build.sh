#!/bin/bash

#builf WebUI
#cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/web
#go install
#cp ~/go/bin/web  ~/go/bin/video_server_web/web
#cp -R ~/go/src/github.com/noChaos1012/noChaos/Server-Video/templates ~/go/bin/video_server_web/

#!/bin/bash


kill -9 $(pgrep api)
kill -9 $(pgrep scheduler)
kill -9 $(pgrep streamserver)
kill -9 $(pgrep web)

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/api
export GO111MODULE=off
go build -o ../bin/api

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/scheduler
 export GO111MODULE=on
 go build -o ../bin/scheduler

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/streamserver
 export GO111MODULE=on
go build -o ../bin/streamserver

cd ~/go/src/github.com/noChaos1012/noChaos/Server-Video/web
 export GO111MODULE=off
go build -o ../bin/web

echo "build finished"
