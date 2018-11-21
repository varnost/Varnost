#!/bin/bash

#python /usr/share/log-gen/apache-fake-log-gen.py -o LOG -p /var/log/apache/ -n 120 -s 1 &

#PID=$!

#sed -i "s/PLACEHOLDER_STRING_HOST/$1/g" /etc/filebeat/filebeat.yml

service filebeat start

#while kill -0 "$PID"; do
  #echo "Waiting for script to DIE"
  #sleep 5
#done
