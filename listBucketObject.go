package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
)

func listBucketObject(minioClient *minio.Client, argsRaw []string) {
	ctx, cancel := context.WithCancel(context.Background())
	var bucketName, subFolder string

	// check second arguments
	if len(argsRaw) > 1 {
		splitedString := strings.Split(argsRaw[1], `/`)
		if len(splitedString) == 1 {
			subFolder = ""
			bucketName = argsRaw[1]
		} else {
			bucketName = splitedString[0]
			subFolder = argsRaw[1][strings.LastIndex(argsRaw[1], `/`)+1:]
		}
	} else {
		panic("Bucket name not found in params")
	}

	defer cancel()

	// list object
	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: subFolder + "/",
	})

	fmt.Printf("Objects list from bucket '%s':\n", bucketName)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Printf("- %s\n", object.Key)
	}
}
