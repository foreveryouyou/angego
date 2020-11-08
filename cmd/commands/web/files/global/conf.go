package global

import "time"

// 站点信息配置
type SiteConfig struct {
	IsProd bool         `yaml:"IsProd,omitempty"`  //
	Port   string       `yaml:"Port,omitempty"`    // web端口
	WebUrl string       `yaml:"WebUrl,omitempty"`  // 网站url
	ResUrl string       `yaml:"ResUrl,omitempty"`  // 静态资源url
	Mongo  MongoConf    `yaml:"MongoDB,omitempty"` // MongoDB 配置
	Redis  RedisConfig  `yaml:"Redis,omitempty"`   // Redis 配置
	WeChat WeChatConfig `yaml:"WeChat,omitempty"`  // 微信 配置
	AliOss aliOss       `yaml:"AliOss,omitempty"`  // 阿里云 配置
}

// Redis 配置
type RedisConfig struct {
	Host      string `yaml:"host,omitempty"`
	Port      string `yaml:"port,omitempty"`
	Database  int    `yaml:"database,omitempty"`
	Password  string `yaml:"password,omitempty"`
	MaxIdle   int    `yaml:"maxIdle,omitempty"`
	MaxActive int    `yaml:"maxActive,omitempty"`
}

// 微信 配置
type WeChatConfig struct {
	AppID          string `yaml:"AppID,omitempty"`
	AppSecret      string `yaml:"AppSecret,omitempty"`
	Token          string `yaml:"Token,omitempty"`
	EncodingAESKey string `yaml:"EncodingAESKey,omitempty"`
}

// MongoDB 配置
type MongoConf struct {
	Host     []string      `yaml:"host,omitempty"`
	Database string        `yaml:"database,omitempty"`
	User     string        `yaml:"user,omitempty"`
	Password string        `yaml:"password,omitempty"`
	Timeout  time.Duration `yaml:"timeout,omitempty"`
}

// 阿里云oss配置
type aliOss struct {
	AccessKeyId     string `yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `yaml:"accessKeySecret,omitempty"`
	Host            string `yaml:"host,omitempty"`
	Endpoint        string `yaml:"endpoint,omitempty"`
	BucketName      string `yaml:"bucketName,omitempty"`
}
