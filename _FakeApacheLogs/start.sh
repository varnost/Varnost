#!/bin/bash

#docker build -t fakelogbeat .

kubectl port-forward logstash-0 5044 &

PID=$!
echo $PID

docker run --rm -it --network host fakelogbeat:latest host.docker.internal

kill $PID
