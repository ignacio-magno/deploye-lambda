package lambd

import (
	"context"
	readfiles "deploye-lambda/read_files"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var client *lambda.Client

func init() {
	client = newClientLambda()
}

func newClientLambda() *lambda.Client {
	// sdk confinrations
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	svc := lambda.NewFromConfig(cfg)

	return svc
}

func GetArnLambdaFunction() string {
	arn, err := client.GetFunctionConfiguration(context.Background(), &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
	})
	if err != nil {
		panic(err)
	}
	return *arn.FunctionArn
}
