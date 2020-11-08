# {{.AppName}}

createdAt: {{.CreatedAt}}

### 开发环境使用
- 复制 `conf.example.yml` 为 `conf.yml` 并根据需求修改其中配置；
- 执行 `dev_build_docker.sh` 创建开发环境docker容器；
- 容器创建成功后执行 `docker exec -it {{.DockerNameDev}} sh` 进入容器交互界面, 然后执行 `go run main.go` 启动项目;

### 生成环境构建
- 复制 `conf.example.yml` 为 `conf.yml` 并根据需求修改其中配置为生产环境配置；
- 执行 `prod_build_docker.sh` 创建生产环境docker容器，服务会自动运行；
