package iamrole

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

var client *iam.Client

func init() {
	client = newClientIAM()
}

func newClientIAM() *iam.Client {
	// sdk confinrations
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	svc := iam.NewFromConfig(cfg)

	return svc
}
