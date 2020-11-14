package upload

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"strings"
	"time"
)

type OssConf struct {
	AccessKeyId     string // AccessKeyId
	AccessKeySecret string // AccessKeySecret
	Host            string // host的格式为 bucketName.endpoint: http://<bucketName>.oss-cn-shanghai.aliyuncs.com
	CallbackUrl     string // 上传回调服务器的URL
	UploadDir       string // 用户上传文件时指定的前缀
	ExpiresIn       int64  // 过期时间, 默认30s
	SizeMin         int64  // 文件大小最小值
	SizeMax         int64  // 文件大小最大值
}

type ossObj struct {
	OssConf
}

func New(conf OssConf) (o *ossObj) {
	o = &ossObj{OssConf: conf}
	return
}

type ConfigStruct struct {
	Expiration string          `json:"expiration"`
	Conditions [][]interface{} `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessId"`           //
	Host        string `json:"host"`               //
	Expire      int64  `json:"expire"`             //
	Signature   string `json:"signature"`          //
	Policy      string `json:"policy"`             //
	Directory   string `json:"dir"`                //
	Callback    string `json:"callback,omitempty"` //
	SizeMin     int64  `json:"sizeMin"`            // 文件大小最小值
	SizeMax     int64  `json:"sizeMax"`            // 文件大小最大值
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

func checkConf(obj *ossObj) error {
	obj.AccessKeyId = strings.TrimSpace(obj.AccessKeyId)
	obj.AccessKeySecret = strings.TrimSpace(obj.AccessKeySecret)
	obj.Host = strings.TrimSpace(obj.Host)
	obj.UploadDir = strings.TrimSpace(obj.UploadDir)
	obj.CallbackUrl = strings.TrimSpace(obj.CallbackUrl)
	if obj.AccessKeyId == "" {
		return errors.New("缺少 AccessKeyId")
	}
	if obj.AccessKeySecret == "" {
		return errors.New("缺少 AccessKeySecret")
	}
	if obj.Host == "" {
		return errors.New("缺少上传 host")
	}
	if obj.ExpiresIn <= 0 {
		obj.ExpiresIn = 30
	}
	return nil
}

func (o *ossObj) GetPolicyToken() (token PolicyToken, err error) {
	err = checkConf(o)
	if err != nil {
		return
	}

	now := time.Now().Unix()
	expiredAt := now + o.ExpiresIn
	var tokenExpire = getGmtIso8601(o.ExpiresIn)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []interface{}
	condition = append(condition, "starts-with", "$key", o.UploadDir)
	config.Conditions = append(config.Conditions, condition)
	if o.SizeMax > 0 {
		config.Conditions = append(config.Conditions, []interface{}{"content-length-range", o.SizeMin, o.SizeMax})
	}

	//calculate signature
	result, err := json.Marshal(config)
	if err != nil {
		return
	}
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(o.AccessKeySecret))
	_, err = io.WriteString(h, deByte)
	if err != nil {
		return
	}
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackBase64 string
	if o.CallbackUrl != "" {
		var callbackParam CallbackParam
		callbackParam.CallbackUrl = o.CallbackUrl
		callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
		callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
		callbackStr, err1 := json.Marshal(callbackParam)
		if err1 != nil {
			err = errors.New(fmt.Sprintf("callback json err: %s", err1))
			return
		}
		callbackBase64 = base64.StdEncoding.EncodeToString(callbackStr)
	}

	token.AccessKeyId = o.AccessKeyId
	token.Host = o.Host
	token.Expire = expiredAt
	token.Signature = string(signedStr)
	token.Directory = o.UploadDir
	token.Policy = string(deByte)
	token.SizeMin = o.SizeMin
	token.SizeMax = o.SizeMax
	if callbackBase64 != "" {
		token.Callback = string(callbackBase64)
	}
	return
}

func getGmtIso8601(expiresIn int64) string {
	now := time.Now()
	now = now.Add(time.Duration(expiresIn * 1e9))
	var tokenExpire = now.UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
