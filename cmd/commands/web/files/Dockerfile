### 构建生产环境
## 构建golang后端服务
#FROM golang:alpine AS build-golang
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN apk add -U git
FROM exidakeng/golang:alpine AS build-golang
## goCache 复用
WORKDIR /app
## 复用 go.mod 缓存
COPY go.mod go.sum ./
RUN go mod download
#RUN go get -u github.com/swaggo/swag/cmd/swag

### 复制应用全部文件
COPY . /app
#RUN swag init
### LDFlags会被`docker build --rm --build-arg`中的LDFlags替换
ARG LDFlags="-s -w"
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFlags" -o app_server main.go

## 构建生产容器
#FROM alpine
#RUN apk add -U tzdata
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
#FROM exidakeng/alpine
FROM scratch
COPY --from=build-golang /etc/localtime /etc/localtime
COPY --from=build-golang /app /app
WORKDIR /app
ENTRYPOINT [ "/app/app_server" ]
