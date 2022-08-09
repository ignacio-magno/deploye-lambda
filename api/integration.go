package api

import (
	"context"
	"deploye-lambda/lambd"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

type Integration struct {
	HttpMethod       string
	Type             types.IntegrationType
	RequestTemplates map[string]string
}

// create integration request
func NewIntegrationRequest(m *MethodToCreate) *Integration {
	var i Integration
	i.HttpMethod = m.HttpMethod

	// consulte type of integration
	fmt.Println("\nSet type of integration")
	fmt.Println("1) AWS_PROXY")
	fmt.Println("2) AWS")
	fmt.Println("3) HTTP")
	fmt.Println("4) MOCK")
	fmt.Println("5) HTTP LAMBDA")

	// read value from console and assign to method
	var opt string
	fmt.Scanln(&opt)
	switch opt {
	case "(1)":
		i.Type = types.IntegrationTypeAwsProxy
	case "2":
		i.Type = types.IntegrationTypeAws
	case "3":
		i.Type = types.IntegrationTypeHttp
	case "4":
		i.Type = types.IntegrationTypeMock
	case "5":
		i.Type = types.IntegrationTypeHttpProxy
	default:
		i.Type = types.IntegrationTypeAwsProxy
	}

	i.RequestTemplates = map[string]string{
		"application/json": "{\"statusCode\": 200}",
	}

	return &i
}

// create integration request
func (a *PathApi) CreateIntegrationRequest(i *Integration) error {
	switch i.Type {
	case types.IntegrationTypeAwsProxy:
		return a.createIntegrationAwsProxy(i)
	case types.IntegrationTypeAws:
		// print option no implemented
		fmt.Println("\nOption no implemented")
		return nil
	case types.IntegrationTypeHttp:
		// print option no implemented
		fmt.Println("\nOption no implemented")
		return nil
	case types.IntegrationTypeMock:
		return a.createIntegrationMock(i)
	case types.IntegrationTypeHttpProxy:
		// print option no implemented
		fmt.Println("\nOption no implemented")
		return nil
	}

	return fmt.Errorf("no type integration exist")
}

func (a *PathApi) createIntegrationAwsProxy(i *Integration) error {

	// print creating integration aws proxy
	fmt.Println("\nCreating integration %v\n", i.Type)
	fmt.Printf("i.HttpMethod: %v\n", i.HttpMethod)

	uri := "arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/" + lambd.GetArnLambdaFunction() + "/invocations"
	fmt.Printf("uri: %v\n", uri)
	// deploy integration type lambda proxy
	_, err := client.PutIntegration(context.Background(), &apigateway.PutIntegrationInput{
		HttpMethod:            aws.String(i.HttpMethod),
		ResourceId:            aws.String(a.id),
		Type:                  i.Type,
		RestApiId:             aws.String(readfiles.ApiId),
		IntegrationHttpMethod: aws.String("POST"),
		Uri:                   aws.String(uri),
		TimeoutInMillis:       aws.Int32(3000),
	})

	if err != nil {
		return err
	}

	// print integration created
	fmt.Println("\nIntegration created")
	return nil
}
