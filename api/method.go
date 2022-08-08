package api

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

// create method options for integration request
type MethodToCreate struct {
	HttpMethod        string
	AuthorizationType string
}

// crete methodtocreate with values recivied from console
func CreateMethod() *MethodToCreate {
	var m MethodToCreate

	// print creating method readfiles.Method
	fmt.Printf("Creating method: %v \n", readfiles.Method)

	// read value from console and assign to method
	m.HttpMethod = readfiles.Method

	// print in blue set authorization type
	fmt.Println("\nSet authorization type")
	fmt.Println("(1) NONE")
	fmt.Println("2) AWS_IAM")
	fmt.Println("3) CUSTOM")
	fmt.Println("4) JWT")

	// read value from console and assign to method
	fmt.Scanln(&m.AuthorizationType)
	switch m.AuthorizationType {
	case "1":
		m.AuthorizationType = "NONE"
	case "2":
		m.AuthorizationType = "AWS_IAM"
	case "3":
		m.AuthorizationType = "CUSTOM"
	case "4":
		m.AuthorizationType = "JWT"
	default:
		m.AuthorizationType = "NONE"
	}
	return &m
}

func (p *PathApi) DeployMethod(m *MethodToCreate) {
	// check if method already exist
	methods := p.getMethod()
	fmt.Printf("methods: %v\n", methods)
	existMethod := false

	for k := range methods {
		if k == m.HttpMethod {
			existMethod = p.methodAlreadyExist(m)
		}
	}

	if !existMethod {
		// print deplying method
		fmt.Printf("Deploying method %v\n", m.HttpMethod)
		// create method
		p.createMethod(m)
	} else {
		// print method is created but no is updated
		fmt.Printf("Method %v is created but no is updated\n", m.HttpMethod)

		// print continuing
		fmt.Println("done method")
	}

}

// method already exist, return if is deleted
func (p *PathApi) methodAlreadyExist(m *MethodToCreate) bool {
	// print method already exist
	fmt.Printf("Method %v already exist\n", m.HttpMethod)
	// consulte if want to update method
	fmt.Println("select option to proced")
	fmt.Println("1) Delete method")
	fmt.Println("2) Continue")

	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		p.deleteMethod(m)
		return true
	case "2":
		return true
	default:
		fmt.Println("Invalid option")
		return p.methodAlreadyExist(m)
	}
}

// delete method
func (p *PathApi) deleteMethod(m *MethodToCreate) {

	_, err := client.DeleteMethod(context.Background(), &apigateway.DeleteMethodInput{
		HttpMethod: aws.String(m.HttpMethod),
		ResourceId: aws.String(p.id),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Method %v deleted\n", m.HttpMethod)
}

// create method
func (p *PathApi) createMethod(m *MethodToCreate) {

	_, err := client.PutMethod(context.Background(), &apigateway.PutMethodInput{
		AuthorizationType: aws.String(m.AuthorizationType),
		HttpMethod:        aws.String(m.HttpMethod),
		ResourceId:        aws.String(p.id),
		RestApiId:         aws.String(readfiles.ApiId),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Method %v created\n", m.HttpMethod)
}
