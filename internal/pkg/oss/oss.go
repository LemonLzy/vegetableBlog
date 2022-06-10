package oss

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

var ossClient = genClient()

func getCosID() string {
	SecretID, ok := os.LookupEnv("COS_SecretID")
	if !ok {
		fmt.Println("COS_SecretID not set")
	}
	fmt.Println(SecretID)
	return SecretID
}

func getCosKey() string {
	SecretKey, ok := os.LookupEnv("COS_SecretKey")
	if !ok {
		fmt.Println("COS_SecretKey not set")
	}
	fmt.Println(SecretKey)
	return SecretKey
}

func genClient() *cos.Client {
	// https://cloud.tencent.com/document/product/436/31215
	u, _ := url.Parse("https://blog-1300597227.cos.ap-guangzhou.myqcloud.com")
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  getCosID(),
			SecretKey: getCosKey(),
		},
	})
	return client
}

func ossGet() {
	result, _, _ := ossClient.Bucket.Get(context.Background(), nil)
	fmt.Println(result)

	//name := "./test/objectPut.go"
	//f := strings.NewReader("test")
	//put, _ := ossClient.Object.Put(context.Background(), name, f, nil)
	//fmt.Println(put)
}
