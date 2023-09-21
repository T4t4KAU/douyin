package minio

import (
	"bytes"
	"context"
	"douyin/pkg/constants"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"net/url"
	"time"
)

var (
	c   *minio.Client
	err error
)

func Init() {
	ctx := context.Background()
	c, err = minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyID, constants.MinioAccessSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	CreateBucket(ctx, constants.MinioVideoBucketName)
	CreateBucket(ctx, constants.MinioImageBucketName)
	CreateBucket(ctx, constants.MinioAvatarBucketName)
	CreateBucket(ctx, constants.MinioBackgroundBucketName)
}

// CreateBucket 创建存储桶
func CreateBucket(ctx context.Context, bucketName string) error {
	options := minio.MakeBucketOptions{}
	if ok, _ := c.BucketExists(ctx, bucketName); !ok {
		err = c.MakeBucket(ctx, bucketName, options)
		if err != nil {
			return err
		}
	}
	return nil
}

// PutToBucketByFile 通过文件传输到OSS
func PutToBucketByFile(ctx context.Context, bucketName string, file *multipart.FileHeader) (minio.UploadInfo, error) {
	f, _ := file.Open()
	defer f.Close()

	options := minio.PutObjectOptions{}
	return c.PutObject(ctx, bucketName, file.Filename, f, file.Size, options)
}

// PutToBucketByBuff 通过缓冲传输将OSS
func PutToBucketByBuff(ctx context.Context, bucketName, filename string, buff *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = c.PutObject(ctx, bucketName, filename, buff, int64(buff.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByPath 通过文件路径传输到OSS
func PutToBucketByPath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = c.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}

func GetObjectURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24 * 7 // 过期时间十年
	reqParams := make(url.Values)
	u, err = c.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}
