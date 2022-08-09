package main

import (
	"fmt"

	// import api gateway aws
	apir "deploye-lambda/api"
	"deploye-lambda/lambd"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

const apiID = "m0gqzb21qk"

type api struct {
	client *apigateway.Client
}

func main() {

	// print select to option to create
	for {
		fmt.Println("1. Create API Gateway")
		fmt.Println("2. Create Cors Method")
		fmt.Println("3. Create lambda function")
		fmt.Println("* exit")

		var options string
		fmt.Scanln(&options)

		switch options {
		case "1":
			method := apir.CreateMethod()
			path := apir.NewPathApi()

			// create all paths need and obtain the id
			path.SetId()

			// print deployming method in blue
			fmt.Printf("\033[1;34m%s\033[0m\n", "Deploying method")
			path.DeployMethod(method)

			// print deployming integration in blue
			fmt.Printf("\033[1;34m%s\033[0m\n", "Deploying integration")
			integration := apir.NewIntegrationRequest(method)
			err := path.CreateIntegrationRequest(integration)
			if err != nil {
				panic(err)
			}
		case "2":
			method := apir.CreateMethod("OPTIONS")
			path := apir.NewPathApi()

			// create all paths need and obtain the id
			path.SetId()

			// print deployming method in blue
			fmt.Printf("\033[1;34m%s\033[0m\n", "Deploying method")
			path.DeployMethod(method)

			// print deployming integration in blue
			fmt.Printf("\033[1;34m%s\033[0m\n", "Deploying integration")
			integration := apir.NewIntegrationRequest(method)
			err := path.CreateIntegrationRequest(integration)
			if err != nil {
				panic(err)
			}
		case "3":
			lambd.DeployLambdaFunction()
		default:
			// print god job
			fmt.Println("\033[1;32m%s\033[0m\n", "God job")
			return
		}
	}

}
