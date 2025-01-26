#!/bin/bash
RUN_NAME=hertz_service
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
# 设置开启go mod
#go env -w GO111MODULE=auto
# 设置go代理
#go env -w GOPROXY=https://goproxy.cn
go build -o output/bin/${RUN_NAME}