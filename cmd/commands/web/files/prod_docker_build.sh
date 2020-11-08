#!/bin/bash
set -x
# 获取当前时间
BuildTime=$(date +'%Y.%m.%d.%H%M%S')
# 将以上变量序列化至 LDFlags 变量中
LDFlags="-s -w
    -X '{{.ModuleName}}/global.BuildTime=${BuildTime}'
"
docker build --rm --build-arg LDFlags="${LDFlags}" -t {{.DockerNameProd}}:latest .
docker-compose -f "prod_docker-compose.yml" up -d --build
