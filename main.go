package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	service "github.com/rexliu0715/go-aws/services"
)

func NewService(options config.LoadOptions) (*service.Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(options.Region))
	if err != nil {
		return nil, err
	}
	return &service.Service{
		Config: cfg,
	}, nil
}
