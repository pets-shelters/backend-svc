package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
)

func (p *Provider) UploadFile(ctx context.Context, body io.Reader, bucket string, key string, contentType string) error {
	_, err := p.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return errors.Wrap(err, "failed to put object to bucket")
	}

	return nil
}

func (p *Provider) DeleteFile(ctx context.Context, bucket string, key string) error {
	_, err := p.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return errors.Wrap(err, "failed to delete object from bucket")
	}

	return nil
}
