#!/bin/bash

cp -R ./templates ./bin/
mkdir ./bin/videos

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


cd bin
nohup ./api > api.outfile 2>&1&
nohup ./scheduler > scheduler.outfile 2>&1&
nohup ./streamserver > streamserver.outfile 2>&1&
nohup ./web > web.outfile 2>&1&

echo "deloy finished"
