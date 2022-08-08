package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

var client *apigateway.Client

func init() {
	client = newClientApiGateway()
}

func newClientApiGateway() *apigateway.Client {
	// sdk confinrations
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	svc := apigateway.NewFromConfig(cfg)

	return svc
}
