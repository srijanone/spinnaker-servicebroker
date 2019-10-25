#!/usr/bin/env bash

spinnaker-servicebroker \
  -logtostderr \
  -GateUrl=${GATE_URL:=http://localhost:8084} \
  --insecure