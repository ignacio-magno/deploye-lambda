package api

import (
	"deploye-lambda/account"
	"deploye-lambda/lambd"
	readfiles "deploye-lambda/read_files"
	"fmt"
)

func (p *PathApi) PutPoliciToLambdaFunction() {
	// print want set policies to invoke lambda function the method
	fmt.Println("\nWant set policies to invoke lambda function the method")
	fmt.Println("(1) Yes")
	fmt.Println("2) No")

	arn := "arn:aws:execute-api:us-west-2:" + account.GetAccountId() + ":" + readfiles.ApiId + "/*/" + readfiles.Method + "/" + p.completePath

	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		pol := lambd.CreatePoliciToApiGateway(arn, p.completePath)
		pol.SetPolicies()
	case "2":
		fmt.Println("done method")
	default:
		fmt.Println("Invalid option")
	}

}
