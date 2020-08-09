#!/bin/bash

cp -R ./templates ./bin/
mkdir ./bin/videos

sh stop.sh

cd bin
nohup ./api > api.outfile 2>&1&
nohup ./scheduler > scheduler.outfile 2>&1&
nohup ./streamserver > streamserver.outfile 2>&1&
nohup ./web > web.outfile 2>&1&

echo "deloy finished"
