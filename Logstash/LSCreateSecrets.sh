#!/bin/bash
kubectl create secret generic ls-ssl \
  --from-file=../../SSL/logstash.crt \
  --from-file=../../SSL/logstash.key \
  --from-file=../../SSL/ca.crt