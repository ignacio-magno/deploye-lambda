package api

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

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

func (a *PathApi) UpdateIntegrationOptions(i *Integration) error {

	var answer string

	// consulte allow headers
	fmt.Println("set allow headers (y/N) \n 'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'")
	fmt.Scanln(&answer)

	if answer == "y" {
		a.responsePatameters["method.response.header.Access-Control-Allow-Headers"] = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'"
	}

	// consulte allow methods
	fmt.Println("set allow methods (y/N) \n 'OPTIONS,GET,PUT,POST,DELETE,PATCH'")
	fmt.Scanln(&answer)

	if answer == "y" {
		a.responsePatameters["method.response.header.Access-Control-Allow-Methods"] = "'OPTIONS,GET,PUT,POST,DELETE,PATCH'"
	}

	// consulte allow origin
	fmt.Println("set allow origin (y/N)")
	fmt.Scanln(&answer)

	if answer == "y" {
		var origin string
		fmt.Println("set allow origin or default *")
		fmt.Scanln(&origin)

		if origin == "" {
			a.responsePatameters["method.response.header.Access-Control-Allow-Origin"] = "'*'"
		} else {
			a.responsePatameters["method.response.header.Access-Control-Allow-Origin"] = "'" + origin + "'"
		}
	}

	// consulte allow credentials
	fmt.Println("set allow credentials (y/N)")
	fmt.Scanln(&answer)

	if answer == "y" {
		a.responsePatameters["method.response.header.Access-Control-Allow-Credentials"] = "'true'"
	}

	//
	a.putResponseMethod(i)
	a.putIntegrationResponse(i)
	return nil
}

func (a *PathApi) getResponseParamters() map[string]bool {
	d := make(map[string]bool)
	for k, _ := range a.responsePatameters {
		d[k] = true
	}

	return d
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
		ResponseParameters: a.getResponseParamters(),
		StatusCode:         aws.String("200"),
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
		HttpMethod:         aws.String(i.HttpMethod),
		ResourceId:         aws.String(a.id),
		RestApiId:          aws.String(readfiles.ApiId),
		StatusCode:         aws.String("200"),
		ResponseParameters: a.responsePatameters,
	})

	if err != nil {
		panic(err)
	}
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
}
