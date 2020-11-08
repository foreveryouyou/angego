#!/usr/bin/env sh
# -o 生成文件名，默认 bindata.go
# -pkg 生成文件的包名，默认 main
# -prefix 生成的文件待去掉的前缀，设置后可以去掉该前缀
go-bindata -o bindata.go -pkg=web -prefix "files/" ./files/...
