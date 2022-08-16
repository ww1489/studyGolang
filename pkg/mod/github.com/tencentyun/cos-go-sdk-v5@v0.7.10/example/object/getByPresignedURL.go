package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

func main() {
	ak := os.Getenv("COS_SECRETID")
	sk := os.Getenv("COS_SECRETKEY")
	u, _ := url.Parse("https://test-1253846586.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ak,
			SecretKey: sk,
			Expire:    time.Hour,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test"
	ctx := context.Background()

	// Normal header way to get object
	resp, err := c.Object.Get(ctx, name, nil)
	log_status(err)
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Get presigned
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, name, ak, sk, time.Hour, nil)
	log_status(err)

	// Get object by presinged url
	resp2, err := http.Get(presignedURL.String())
	log_status(err)
	bs2, _ := ioutil.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("result2 is : %s\n", string(bs2))

	fmt.Printf("%v\n\n", bytes.Compare(bs2, bs) == 0)

}
