package oss

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

const (
	FileDir = "blog/"
)

var (
	cosURL    string
	secretID  string
	secretKey string
	ossClient *cos.Client
)

func init() {
	cosURL = getCosURL()
	secretID = getCosID()
	secretKey = getCosKey()
	ossClient = genClient()
}

func getCosID() string {
	id, ok := os.LookupEnv("COS_SecretID")
	if !ok {
		fmt.Println("COS_SecretID not set")
	}
	return id
}

func getCosKey() string {
	key, ok := os.LookupEnv("COS_SecretKey")
	if !ok {
		fmt.Println("COS_SecretKey not set")
	}
	return key
}

func getCosURL() string {
	link, ok := os.LookupEnv("COS_URL")
	if !ok {
		fmt.Println("COS_URL not set")
	}
	return link
}

// genClient 初始化COS对象
func genClient() *cos.Client {
	// https://cloud.tencent.com/document/product/436/31215
	u, _ := url.Parse(cosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
	return client
}

// Put 上传文件至COS
func Put(fileHeader *multipart.FileHeader) string {
	filename := FileDir + fileHeader.Filename
	file, _ := fileHeader.Open()
	_, err := ossClient.Object.Put(context.Background(), filename, file, nil)
	if err != nil {
		fmt.Println(err)
	}
	return cosURL + "/" + filename
}
