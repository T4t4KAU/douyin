package oss

import (
	"bytes"
	"context"
	"douyin/cmd/publish/pkg"
	"douyin/pkg/constants"
	"douyin/pkg/minio"
)

func Init() {
	minio.Init()
}

const ChunkSize = 1024 * 128

func UploadVideo(data []byte, filename string) (bool, error) {
	b := bytes.NewBuffer(data)
	_, err := minio.PutToBucketByBuff(context.Background(),
		constants.MinioVideoBucketName, filename, b)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UploadImage(data []byte, filename string) (bool, error) {
	b := bytes.NewBuffer(data)
	_, err := minio.PutToBucketByBuff(context.Background(),
		constants.MinioImageBucketName, filename, b)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UploadFile(f pkg.FilePair) {
	go UploadVideo(f.VideoData, f.VideoName)
	go UploadImage(f.ImageData, f.ImageName)
}
