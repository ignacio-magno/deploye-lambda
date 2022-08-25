package lambd

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func UpdateFunctionLambdaCode() error {

	code, err := readfiles.DeployCode()
	if err != nil {
		panic(err)
	}

	resp, err := client.UpdateFunctionCode(context.Background(), &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
		ZipFile:      code,
	})

	fmt.Printf("resp.State: %v\n", resp.State)

	return err
}
