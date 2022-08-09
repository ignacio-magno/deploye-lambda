package lambd

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

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
	fname := readfiles.GetNameFunctionLambda()
	fmt.Printf("fname: %v\n", fname)
	arn, err := client.GetFunctionConfiguration(context.Background(), &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(fname),
	})
	if err != nil {
		panic(err)
	}
	return *arn.FunctionArn
}
