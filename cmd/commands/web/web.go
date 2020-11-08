package web

import (
	"fmt"
	"github.com/foreveryouyou/angego/utils"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var ModuleName string     // 项目go_module名: 字母、数字、下划线、点、斜杠
var AppName string        // 项目名: 字母、数字、下划线
var DockerNameDev string  // docker容器名(dev): 字母、数字、下划线，默认golang
var DockerNameProd string // docker容器名(prod): 字母、数字、下划线，默认golang
var ok bool

// 终端颜色
var green = color.FgGreen.Render
var red = color.FgRed.Render

func Web() {
	fmt.Println(green("[项目设置]"))
	// 项目go_module名
	ok = false
	for !ok {
		color.Primary.Println("项目go_module名(字母、数字、下划线、斜杠):")
		fmt.Scanln(&ModuleName)
		if IsValidModuleName(ModuleName) {
			ok = true
		}
	}

	// 项目名
	ok = false
	for !ok {
		color.Primary.Println("项目名(字母、数字、下划线):")
		fmt.Scanln(&AppName)
		if IsValidAppName(AppName) {
			ok = true
		}
	}

	// docker容器名(dev)
	ok = false
	for !ok {
		color.Primary.Println("docker容器名dev(字母、数字、下划线，默认 {AppName}_dev):")
		fmt.Scanln(&DockerNameDev)
		if IsValidDockerName(DockerNameDev) {
			ok = true
		} else {
			DockerNameDev = AppName + "_dev"
			ok = true
		}
	}

	// docker容器名(prod)
	ok = false
	for !ok {
		color.Primary.Println("docker容器名prod(字母、数字、下划线，默认 {AppName}_golang):")
		fmt.Scanln(&DockerNameProd)
		if IsValidDockerName(DockerNameProd) {
			ok = true
		} else {
			DockerNameProd = AppName + "_golang"
			ok = true
		}
	}

	fmt.Println(green("[项目配置]"))
	fmt.Println(green("go_module名:"), ModuleName)
	fmt.Println(green("项目名:"), AppName)
	fmt.Println(green("docker容器名dev:"), DockerNameDev)
	fmt.Println(green("docker容器名prod:"), DockerNameProd)
	fmt.Println()
	fmt.Println(green("[开始创建...]"))

	CreateProject()
}

// 创建项目
func CreateProject() {
	var err error
	if utils.FileExist(AppName) {
		fmt.Println("[", AppName, "]已存在")
		return
	}
	for _, file := range AssetNames() {
		fmt.Println(green("[CreateFile]"), file)
		if err = GenFiles("./"+AppName, file); err != nil {
			fmt.Println(red("[CreateFile err]:"), file)
		}
	}
	fmt.Println(green("[创建成功]"), AppName)
}

//GenFile restores an asset under the given directory
func GenFile(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	// ↓↓↓这里加入自定义处理
	var code string
	if name == "go.mod" ||
		strings.HasSuffix(name, ".go") {
		code = string(data)
		code = strings.ReplaceAll(code, "{{.ModuleName}}", ModuleName)
	}
	if name == "dev_docker-compose.yml" {
		code = string(data)
		code = strings.ReplaceAll(code, "{{.DockerNameDev}}", DockerNameDev)
	}
	if name == "prod_docker-compose.yml" {
		code = string(data)
		code = strings.ReplaceAll(code, "{{.DockerNameProd}}", DockerNameProd)
	}
	if name == "prod_docker_build.sh" {
		code = string(data)
		code = strings.ReplaceAll(code, "{{.ModuleName}}", ModuleName)
		code = strings.ReplaceAll(code, "{{.DockerNameProd}}", DockerNameProd)
	}
	if name == "README.md" {
		code = string(data)
		code = strings.ReplaceAll(code, "{{.AppName}}", AppName)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		code = strings.ReplaceAll(code, "{{.CreatedAt}}", createdAt)
	}
	if code != "" {
		data = []byte(code)
	}
	// ↑↑↑这里加入自定义处理
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

//GenFiles restores an asset under the given directory recursively
func GenFiles(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return GenFile(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

//字母、数字、下划线、点、斜杠，字母开头、字母结尾
func IsValidModuleName(name string) bool {
	reg := `^[a-z][a-z0-9_/\.]*[a-z0-9]$`
	return regexp.MustCompile(reg).MatchString(name)
}

//字母、数字、下划线，字母开头
func IsValidAppName(name string) bool {
	reg := `^[A-Za-z](\w)*[A-Za-z0-9]$`
	return regexp.MustCompile(reg).MatchString(name)
}

//字母、数字、下划线，字母开头
func IsValidDockerName(name string) bool {
	reg := `^[a-z][a-z0-9_]*[a-z0-9]$`
	return regexp.MustCompile(reg).MatchString(name)
}
