#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o dev
rsync -az --delete --progress /home/tonnn/go/src/zonstfe_api/developer/dev   -e 'ssh ' tonnn@111.231.137.127:/home/tonnn/zonstfe_api/
rm dev
rm developer