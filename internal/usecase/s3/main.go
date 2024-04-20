package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pkg/errors"
)

type Provider struct {
	client *s3.Client
}

func NewProvider(cfg configs.S3) (*Provider, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if cfg.WriteEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           cfg.WriteEndpoint,
				SigningRegion: cfg.Region,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	creds := credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")

	awsCfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(creds),
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws config")
	}

	awsS3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &Provider{
		awsS3Client,
	}, nil
}
