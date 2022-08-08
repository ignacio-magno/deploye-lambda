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

	// deploy integration type lambda proxy
	_, err := client.PutIntegration(context.Background(), &apigateway.PutIntegrationInput{
		HttpMethod:            aws.String(i.HttpMethod),
		ResourceId:            aws.String(a.id),
		Type:                  i.Type,
		RestApiId:             aws.String(readfiles.ApiId),
		IntegrationHttpMethod: aws.String("POST"),
		Uri:                   aws.String("arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/" + lambd.GetArnLambdaFunction() + "/invocations"),
		TimeoutInMillis:       aws.Int32(3000),
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *PathApi) createIntegrationMock(i *Integration) error {
	// print creating integration mock
	fmt.Println("\nCreating integration mock")
	fmt.Printf("a.id: %v\n", a.id)
	_, err := client.PutIntegration(context.Background(), &apigateway.PutIntegrationInput{
		HttpMethod:       aws.String(i.HttpMethod),
		ResourceId:       aws.String(a.id),
		Type:             i.Type,
		RestApiId:        aws.String(readfiles.ApiId),
		RequestTemplates: i.RequestTemplates,
	})

	if err != nil {
		return err
	}

	// print how is integration mock, requiere method reponse
	a.putResponseMethod(i)
	a.putIntegrationResponse(i)
	return nil
}

// put response method
func (a *PathApi) putResponseMethod(i *Integration) {

	_, err := client.PutMethodResponse(context.Background(), &apigateway.PutMethodResponseInput{
		HttpMethod: aws.String(i.HttpMethod),
		ResourceId: aws.String(a.id),
		RestApiId:  aws.String(readfiles.ApiId),
		ResponseModels: map[string]string{
			"application/json": "Empty",
		},
		ResponseParameters: map[string]bool{
			"method.response.header.Access-Control-Allow-Origin": true,
		},
		StatusCode: aws.String("200"),
	})
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		fmt.Println("eliminando metodo")
		_, err := client.DeleteMethodResponse(context.Background(), &apigateway.DeleteMethodResponseInput{
			HttpMethod: aws.String("OPTIONS"),
			ResourceId: aws.String(a.id),
			RestApiId:  aws.String(readfiles.ApiId),
			StatusCode: aws.String("200"),
		})
		if err != nil {
			panic(err)
		} else {
			a.putResponseMethod(i)
		}
	}
}

// put integration response
func (a *PathApi) putIntegrationResponse(i *Integration) {
	res, err := client.PutIntegrationResponse(context.Background(), &apigateway.PutIntegrationResponseInput{
		HttpMethod: aws.String(i.HttpMethod),
		ResourceId: aws.String(a.id),
		RestApiId:  aws.String(readfiles.ApiId),
		StatusCode: aws.String("200"),
		ResponseParameters: map[string]string{
			"method.response.header.Access-Control-Allow-Origin": "'*'",
		},
	})

	if err != nil {
		panic(err)
	}
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
}
