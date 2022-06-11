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
	FileDir = "test/"
	CosURL  = "https://blog-1300597227.cos.ap-guangzhou.myqcloud.com/"
)

var ossClient = genClient()

func getCosID() string {
	SecretID, ok := os.LookupEnv("COS_SecretID")
	if !ok {
		fmt.Println("COS_SecretID not set")
	}
	return SecretID
}

func getCosKey() string {
	SecretKey, ok := os.LookupEnv("COS_SecretKey")
	if !ok {
		fmt.Println("COS_SecretKey not set")
	}
	return SecretKey
}

func getCosURL() string {
	CosUrl, ok := os.LookupEnv("COS_URL")
	if !ok {
		fmt.Println("COS_URL not set")
	}
	return CosUrl
}

// genClient 初始化COS对象
func genClient() *cos.Client {
	// https://cloud.tencent.com/document/product/436/31215
	u, _ := url.Parse(CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  getCosID(),
			SecretKey: getCosKey(),
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
	return CosURL + filename
}
