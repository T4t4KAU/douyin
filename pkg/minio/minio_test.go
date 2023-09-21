package minio

import (
	"bytes"
	"context"
	"douyin/pkg/constants"
	"fmt"
	"io/ioutil"

	"testing"
)

func TestBucketExist(t *testing.T) {
	Init()
	ctx := context.Background()
	exists, err := c.BucketExists(ctx, constants.MinioVideoBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("%v found\n", constants.MinioVideoBucketName)
	} else {
		fmt.Println("not found")
	}
}

func TestPutToBucketByBuff(t *testing.T) {
	Init()
	b, err := ioutil.ReadFile("./videos/test1.mp4")
	if err != nil {
		t.Errorf("read error: %v", err)
	}
	buff := bytes.NewBuffer(b)

	info, err := PutToBucketByBuff(context.Background(), constants.MinioVideoBucketName, "test1.mp4", buff)
	if err != nil {
		t.Errorf("put error: %v", err)
	}
	fmt.Printf("%#v\n", info)
}

func TestGetObjectURL(t *testing.T) {
	Init()
	u, err := GetObjectURL(context.Background(), "images", "1000.jpg")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	println(u.String())
}
