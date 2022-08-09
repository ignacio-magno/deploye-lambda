package api

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func Deploy() {
	res, err := client.CreateDeployment(context.Background(), &apigateway.CreateDeploymentInput{
		RestApiId: aws.String(readfiles.ApiId),
		StageName: aws.String("main"),
	})

	if err != nil {
		panic(err)
	}
	fmt.Printf("res.ResultMetadata: %v\n", res.ResultMetadata)
}
