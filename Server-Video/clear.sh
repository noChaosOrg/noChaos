#!/bin/bash

for apiPID in $(pgrep api)
do
    kill -9 $apiPID
done

for schedulerPID in $(pgrep scheduler)
do
  kill -9 $schedulerPID
done

for streamserverPID in $(pgrep streamserver)
do
  kill -9 $streamserverPID
done

for webPID in $(pgrep web)
do
  kill -9 $webPID
done

rm -rf ./bin/api.outfile
rm -rf ./bin/scheduler.outfile
rm -rf ./bin/streamserver.outfile
rm -rf ./bin/web.outfile

rm -rf ./bin/api
rm -rf ./bin/scheduler
rm -rf ./bin/streamserver
rm -rf ./bin/web

echo "clear finished"


