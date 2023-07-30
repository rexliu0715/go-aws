package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/net/context"
)

type Service struct {
	Config aws.Config
}

func (s *Service) GetObjects(ctx context.Context, input s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return s3.NewFromConfig(s.Config).ListObjectsV2(ctx, &input)
}
