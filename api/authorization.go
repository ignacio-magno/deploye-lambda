package api

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

func (a *MethodToCreate) GetIdAuthorization() string {
	resp, err := client.GetAuthorizers(context.Background(), &apigateway.GetAuthorizersInput{
		RestApiId: aws.String(readfiles.ApiId),
	})

	if err != nil {
		panic(err)
	}

	if a.AuthorizationType == "COGNITO_USER_POOLS" {
		fmt.Println("Select authorization type")
		for i, a2 := range resp.Items {
			fmt.Printf("%v) %v\n", i, *a2.Name)
		}

		var answer string
		// read answer from user
		fmt.Scanln(&answer)

		indice, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
			return ""
		}
		id := *resp.Items[indice].Id
		fmt.Printf("id: %v\n", id)
		return id
	}
	return ""
}

func (p *PathApi) SetAuthorization(m *MethodToCreate) {
	id := m.GetIdAuthorization()
	fmt.Printf("id: %v\n", id)
	resp, err := client.UpdateMethod(context.Background(), &apigateway.UpdateMethodInput{
		HttpMethod: aws.String(readfiles.Method),
		ResourceId: aws.String(p.id),
		RestApiId:  aws.String(readfiles.ApiId),
		PatchOperations: []types.PatchOperation{
			{
				Op:    types.OpReplace,
				From:  aws.String("UpdateMethod"),
				Path:  aws.String("/authorizationType"),
				Value: aws.String(m.AuthorizationType),
			},
			{
				Op:    types.OpReplace,
				From:  aws.String("UpdateMethod"),
				Path:  aws.String("/authorizerId"),
				Value: aws.String(id),
			},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("resp.AuthorizerId: %v\n", resp.AuthorizerId)
}
