package lambd

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type Policy struct {
	Action      string
	Principal   string
	StatementId string
	SourceArn   string
}

func CreatePoliciToApiGateway(arnMethodLambda string, path string) *Policy {
	return &Policy{
		Action:      "lambda:InvokeFunction",
		Principal:   "apigateway.amazonaws.com",
		StatementId: strings.ReplaceAll(path, "/", "-"),
		SourceArn:   arnMethodLambda,
	}
}

func (p *Policy) SetPolicies() {
	fmt.Println("Set policies")
	_, err := client.AddPermission(context.Background(), &lambda.AddPermissionInput{
		Action:       aws.String(p.Action),
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
		Principal:    aws.String(p.Principal),
		StatementId:  aws.String(p.StatementId),
		SourceArn:    aws.String(p.SourceArn),
	})

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		p.removePolici()
		p.SetPolicies()
	}

	fmt.Println("Policies set")
}

func (p *Policy) removePolici() bool {
	_, err := client.RemovePermission(context.Background(), &lambda.RemovePermissionInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
		StatementId:  aws.String(p.StatementId),
	})

	panic(err)
}
