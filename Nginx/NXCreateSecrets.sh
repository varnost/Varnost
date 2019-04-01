#!/bin/bash
kubectl create secret generic nx-ssl \
  --from-file=../../SSL/web.crt \
  --from-file=../../SSL/web.key
kubectl create secret generic htpasswd \
  --from-file=../../SSL/.htpasswd