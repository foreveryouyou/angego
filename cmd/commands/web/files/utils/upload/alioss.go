package upload

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path/filepath"
	"time"
)

var (
	ossBucketName      = ""
	ossAccessKeyId     = ""
	ossAccessKeySecret = ""
	ossEndPoint        = ""
)

func InitBucketInfo(bucketName, accessKeyId, accessKeySecret, endPoint string) {
	ossBucketName = bucketName
	ossAccessKeyId = accessKeyId
	ossAccessKeySecret = accessKeySecret
	ossEndPoint = endPoint
}

// GetBucket
func GetBucket(bucketName string) (bucket *oss.Bucket, err error) {
	// New client
	client, err := oss.New(ossEndPoint, ossAccessKeyId, ossAccessKeySecret)
	if err != nil {
		return
	}

	// Create bucket
	/*err = client.CreateBucket(bucketName)
	if err != nil {
		return
	}*/

	// Get bucket
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		return
	}

	return
}

// PutObjectSmall 小文件简单上传
func PutObjectSmall(objectKey string, filePath string) (path string, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	path = objectKey

	// Upload an object with local file name, user need not open the file.
	acl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObjectFromFile(objectKey, filePath, acl)
	return
}

// PutObjectSmallACLPrivate 小文件简单上传 私有
func PutObjectSmallACLPrivate(objectKey string, filePath string) (path string, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	path = objectKey

	// Upload an object with local file name, user need not open the file.
	acl := oss.ObjectACL(oss.ACLPrivate)
	err = bucket.PutObjectFromFile(objectKey, filePath, acl)
	return
}

func GetSignUrl(objectKey string, expiredIn int64) (signedURL string, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	signedURL, err = bucket.SignURL(objectKey, oss.HTTPGet, expiredIn)
	return
}

// IsObjectExist oss 对象是否存在
func IsObjectExist(objectKey string) (isExist bool) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 判断文件是否存在。
	isExist, err = bucket.IsObjectExist(objectKey)
	return
}

// PutObjectBytes 小文件简单上传
func PutObjectBytes(objectKey string, aByte []byte) (path string, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	path = objectKey

	// 上传Byte数组
	acl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObject(objectKey, bytes.NewReader(aByte), acl)
	if err != nil {
		return
	}
	return
}

// GetObjectLastModified 返回object的最后修改时间
func GetObjectLastModified(objectKey string) (lastModified int64, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 获取文件元信息。
	props, err := bucket.GetObjectDetailedMeta(objectKey)
	if err != nil {
		return
	}
	lm, _ := props["Last-Modified"]
	if len(lm) < 1 || lm[0] == "" {
		return
	}
	t, err := time.Parse(time.RFC1123, lm[0])
	if err != nil {
		return
	}
	lastModified = t.Unix()
	return
}

// SetObjectTagging 设置object的tagging
func SetObjectTagging(objectKey string, tags map[string]string) (err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 设置Tagging规则。
	var Tags []oss.Tag
	for k, v := range tags {
		Tags = append(Tags, oss.Tag{
			Key:   k,
			Value: v,
		})
	}
	tagging := oss.Tagging{
		Tags: Tags,
	}
	err = bucket.PutObjectTagging(objectKey, tagging)
	return
}

// GetObjectTagging 获取object的tagging
func GetObjectTagging(objectKey string) (tags []oss.Tag, err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 获取Tagging规则
	result, err := bucket.GetObjectTagging(objectKey)
	if err != nil {
		return
	}
	tags = result.Tags
	return
}

// PutAllDirFiles 将指定目录所有文件上传
func PutAllDirFiles(dir string) (err error) {
	// 上传目录中的文件
	err = filepath.Walk(dir, func(filePath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() || f.Name() == ".DS_Store" {
			return nil
		}
		// 上传文件到OSS
		_, err = PutObjectSmall(filePath, filePath)
		return err
	})
	return
}

// 列举指定前缀（目录）下的文件信息
// lor.IsTruncated: 是否未列举完
func ListFiles(prefixStr, markerStr, delimiter string, maxKeys int) (lor oss.ListObjectsResult, err error) {
	// 获取存储空间。
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 遍历文件
	var opts []oss.Option
	marker := oss.Marker(markerStr)
	prefix := oss.Prefix(prefixStr)
	if maxKeys <= 0 || maxKeys > 200 {
		maxKeys = 100
	}
	opts = append(opts, oss.MaxKeys(maxKeys))
	opts = append(opts, marker)
	opts = append(opts, prefix)
	opts = append(opts, oss.Delimiter(delimiter))
	// 列出文件
	lor, err = bucket.ListObjects(opts...)
	return
}

// 创建目录
func CreateDir(dir string) (err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 判断文件是否存在。
	isExist, err := bucket.IsObjectExist(dir)
	if err != nil {
		return
	}
	if isExist {
		return
	}
	// 上传空Byte数组
	acl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObject(dir, bytes.NewReader([]byte{}), acl)
	return
}

// 删除文件
func DeleteObject(objectKey string) (err error) {
	// Get bucket
	bucket, err := GetBucket(ossBucketName)
	if err != nil {
		return
	}

	// 删除文件
	err = bucket.DeleteObject(objectKey)
	return
}
