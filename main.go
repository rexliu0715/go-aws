package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	service "github.com/rexliu0715/go-aws/services"
)

func NewService(optFns ...func(*config.LoadOptions) error) (*service.Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), optFns...)
	if err != nil {
		return nil, err
	}
	return &service.Service{
		Config: cfg,
	}, nil
}
