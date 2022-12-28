package s3

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3Cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/soa/product-api/internal/config"
)

type Image struct {
}

type S3Client struct {
	api    *s3.Client
	bucket string
}

func NewS3(cfg *config.Config) (*S3Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               cfg.S3.URL,
				SigningRegion:     cfg.S3.Region,
				HostnameImmutable: true,
			}, nil
		},
	)
	credentials := aws.CredentialsProviderFunc(
		func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     cfg.S3.AccessKeyID,
				SecretAccessKey: cfg.S3.SecretAccessKey,
			}, nil
		})

	s3Cfg, err := s3Cfg.LoadDefaultConfig(
		context.Background(),
		s3Cfg.WithEndpointResolverWithOptions(resolver),
		s3Cfg.WithCredentialsProvider(credentials),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(s3Cfg)
	return &S3Client{
		api:    client,
		bucket: cfg.S3.Bucket,
	}, nil
}

func (s *S3Client) PutImage(ctx context.Context, img []byte, filename string) error {
	_, err := s.api.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(img),
	})

	return err
}

func (s *S3Client) GetImage(ctx context.Context, filename string) ([]byte, error) {
	res, err := s.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
