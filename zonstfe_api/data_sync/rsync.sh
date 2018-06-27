#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o data_sync
rsync -az --delete --progress /home/tonnn/go/src/zonstfe_api/data_sync/data_sync   -e 'ssh ' tonnn@111.231.137.127:/home/tonnn/zonstfe_api/
rm data_sync



