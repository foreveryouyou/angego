### 构建生产环境
## 构建golang后端服务
#FROM golang:alpine AS build-golang
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN apk add -U git
FROM exidakeng/golang:alpine AS build-golang
## goCache 复用
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
#RUN go get -u github.com/swaggo/swag/cmd/swag

###
ADD . /go/src/app
WORKDIR /go/src/app
#RUN swag init
### LDFlags会被`docker build --rm --build-arg`中的LDFlags替换
ARG LDFlags="-s -w"
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFlags" -o app_server main.go

## 构建生产容器
#FROM alpine
#RUN apk add -U tzdata
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
FROM exidakeng/alpine
COPY --from=build-golang /go/src/app /app
WORKDIR /app
EXPOSE 80
#CMD [ "sh", "-c", "/app/app_server" ]
