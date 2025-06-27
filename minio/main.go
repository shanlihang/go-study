package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    endpoint := "http://127.0.0.1:9000" // 或者你本地的 IP:PORT
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    useSSL := false

    // 初始化 MinIO 客户端
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("✅ 成功连接 MinIO")
    fmt.Println("MinIO",minioClient)

    
}
