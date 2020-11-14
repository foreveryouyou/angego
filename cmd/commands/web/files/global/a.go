package global

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Global 全局变量或属性
var (
	// 通过 go build -ldflags 的 -X 参数传入
	BuildTime string // build时间
)
var BuildGoVersion = runtime.Version() // build的go版本信息

// 站点配置
var SConf = &SiteConfig{
	IsProd: false,
	Port:   "80",
	Mongo:  MongoConf{},
	Redis:  RedisConfig{},
	WeChat: WeChatConfig{},
	AliOss: aliOss{},
}

// InitConfig 初始化站点配置
func InitConfig(configFile string) {
	err := initConf(configFile, &SConf)
	if err != nil {
		log.Fatal("配置初始化失败:", err)
	}
	return
}

// conf pointer
func initConf(filename string, conf interface{}) (err error) {
	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		return
	}

	// put the file's contents as yaml to the default configuration(conf)
	err = yaml.Unmarshal(data, conf)
	return nil
}
