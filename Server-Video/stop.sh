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

echo "stop finished"

