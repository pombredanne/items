#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o dev_admin
rsync -az --delete --progress   /home/tonnn/go/src/zonstfe_api/developer_admin/dev_admin   -e 'ssh ' tonnn@111.231.137.127:/home/tonnn/zonstfe_api/
rm dev_admin
rm developer_admin