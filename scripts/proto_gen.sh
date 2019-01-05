#!/bin/bash

# Call from scripts/ folder.
PROTO_PATH=../message
protoc --proto_path=$PROTO_PATH \
    --go_out=$PROTO_PATH \
    ${PROTO_PATH}/*.proto
