package utils

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Ctx    context.Context
}

func NewMinioClient(endpoint, accessKey, secretKey string, useSSL bool) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Minio client: %w", err)
	}
	return &MinioClient{
		Client: client,
		Ctx:    context.Background(),
	}, nil
}

func (m *MinioClient) EnsureBucket(bucketName, location string) error {
	exists, err := m.Client.BucketExists(m.Ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.Client.MakeBucket(m.Ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MinioClient) UploadFile(bucket, objectName, filePath, contentType string) error {
	_, err := m.Client.FPutObject(m.Ctx, bucket, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (m *MinioClient) DownloadFile(bucket, objectName, destPath string) error {
	return m.Client.FGetObject(m.Ctx, bucket, objectName, destPath, minio.GetObjectOptions{})
}

func (m *MinioClient) GeneratePresignedURL(bucket, objectName string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	return m.Client.PresignedGetObject(m.Ctx, bucket, objectName, expiry, reqParams)
}

func (m *MinioClient) ListObjects(bucket string, recursive bool) ([]minio.ObjectInfo, error) {
	objectCh := m.Client.ListObjects(m.Ctx, bucket, minio.ListObjectsOptions{Recursive: recursive})
	results := []minio.ObjectInfo{}
	for obj := range objectCh {
		if obj.Err != nil {
			return nil, obj.Err
		}
		results = append(results, obj)
	}
	return results, nil
}

func (m *MinioClient) DeleteObject(bucket, objectName string) error {
	return m.Client.RemoveObject(m.Ctx, bucket, objectName, minio.RemoveObjectOptions{})
}
