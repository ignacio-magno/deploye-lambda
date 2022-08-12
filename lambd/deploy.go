package lambd

import (
	"context"
	"deploye-lambda/iamrole"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func DeployLambdaFunction() {
	if existFunction() {
		// print in blue function already exist, you want delete this
		fmt.Printf("\x1b[34m%s\x1b[0m\n", "Function already exist, you want delete this? (y/N")
		var answer string
		fmt.Scanln(&answer)
		if answer == "y" {
			deleteFunctionLambda()
			// print function deleted
			fmt.Printf("\x1b[32m%s\x1b[0m\n", "Function deleted")
			DeployLambdaFunction()
		}
	} else {

		if arnRol, ok := iamrole.ExistRol(); ok {
			//print rol exist, deploy function
			fmt.Println("Rol exist deploying function....")

			deployFunctionLambda(arnRol)
			fmt.Println("function created")
		} else {
			// print in red role not exist, you want create this
			fmt.Printf("\x1b[31m%s\x1b[0m\n", "Role not exist, you want create this? (y/N")
			var answer string
			fmt.Scanln(&answer)
			if answer == "y" {
				iamrole.CreateRole()
				DeployLambdaFunction()
			}
		}
	}
}

func deployFunctionLambda(arnRole string) {
	fmt.Println("Deploy function lambda")

	code, err := readfiles.DeployCode()
	if err != nil {
		panic(err)
	}

	_, err = client.CreateFunction(context.Background(), &lambda.CreateFunctionInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
		Description:  aws.String("Function lambda"),
		Runtime:      types.RuntimeGo1x,
		Role:         aws.String(arnRole),
		Handler:      aws.String("main"),
		Code: &types.FunctionCode{
			ZipFile: code,
		},
		Environment: &types.Environment{
			Variables: readfiles.GetEnvironmentVariables(),
		},
		Timeout: aws.Int32(30000),
	})
	if err != nil {
		panic(err)
	}
}

func deleteFunctionLambda() {
	fmt.Println("Delete function lambda")

	_, err := client.DeleteFunction(context.Background(), &lambda.DeleteFunctionInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
	})

	if err != nil {
		panic(err)
	}
}

func existFunction() bool {
	_, err := client.GetFunction(context.Background(), &lambda.GetFunctionInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
	})

	// TODO repart this part, after err can be other errror and not not exist function
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return false
	}

	return true

}

func AddEnvironmentVariables() {
	fmt.Printf("readfiles.GetEnvironmentVariables(): %v\n", readfiles.GetEnvironmentVariables())
	_, err := client.UpdateFunctionConfiguration(context.Background(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(readfiles.GetNameFunctionLambda()),
		Environment: &types.Environment{
			Variables: readfiles.GetEnvironmentVariables(),
		},
	})

	if err != nil {
		panic(err)
	}
}
