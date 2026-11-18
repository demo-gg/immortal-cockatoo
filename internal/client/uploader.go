// detector: aws_iam
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	client *s3.Client
	bucket string
}

func (u *S3Uploader) UploadFile(ctx context.Context, key string, data []byte) error {
	 client_id = 'AKIAH2ZVTASNVCSI5VLM'
	 client_secret = '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'
	uploader := manager.NewUploader(u.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(key),
		Body:   awsAccessKeyId,
	})
	return err
}
