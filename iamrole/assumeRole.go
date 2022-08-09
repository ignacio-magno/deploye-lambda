package iamrole

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func CreateRole() {
	for {
		if _, ok := ExistRol(); ok {
			// print exist role
			fmt.Println(" Role exist!")

			// consulte if want delete role and policies
			var answer string
			fmt.Print(" 1) delete all role and policies? ")
			fmt.Print(" 2) put new roles? ")
			fmt.Print(" * exit? ")
			fmt.Scanln(&answer)

			switch answer {
			case "1":
				deleteRoleAndPolicies()
			case "2":
				policies := getPoliciesInRole()
				// print existing policies
				fmt.Println(" Existing policies: ")
				for _, v := range policies {
					fmt.Printf("%v\n", v)
				}

				// print policies in folder
				fmt.Println("select policie to set")
				files, getFile := getRolesInFolder()
				for i, v := range files {
					fmt.Printf("%v) %v\n", i, v)
				}

				var pol string
				fmt.Scanln(&pol)

				// crea un objeto role policies basado en el nombre del archivo seleccionado
				role := getFile(pol)

				// print creating polici
				fmt.Println("creating polici...")
				if role.existPolicieInRol() {
					role.deletePolici()
				} else {
					role.createPolicie()
				}
			default:
				return
			}
		} else {
			createRole()
			break
		}
	}
}

// return if exist rol and arn string
func ExistRol() (string, bool) {
	// find if exist role
	res, err := client.GetRole(context.Background(), &iam.GetRoleInput{
		RoleName: aws.String(readfiles.GetNameFunctionLambda() + "-role"),
	})

	if err != nil {
		fmt.Println("rol no encontrado")
		fmt.Printf("err.Error(): %v\n", err.Error())
		return "", false
	} else {
		fmt.Println("rol encontrado")
		return *res.Role.Arn, true
	}
}

// create role default
func createRole() {

	// print creating role
	fmt.Println(" Creating role...")
	trust := NewTrusPolici()

	// create role
	_, err := client.CreateRole(context.Background(), &iam.CreateRoleInput{
		RoleName:                 aws.String(readfiles.GetNameFunctionLambda() + "-role"),
		AssumeRolePolicyDocument: aws.String(trust.GetBytesFile()),
	})
	if err != nil {
		panic(err)
	}

	// print creating policies
	fmt.Println(" Creating policies logs...")
	policies := NewPutLogPolici()

	// put policies
	policies.putPolicies()

	// print done
	fmt.Println(" Done!")

}

func deleteRoleAndPolicies() {
	// print removing policies
	fmt.Println(" Removing policies from role...")
	deleteAllPolicies()

	// print removing role
	fmt.Println(" Removing role...")
	deleteRol()

}

func deleteRol() {
	_, err := client.DeleteRole(context.Background(), &iam.DeleteRoleInput{
		RoleName: aws.String(readfiles.GetNameFunctionLambda() + "-role"),
	})
	if err != nil {
		panic(err)
	}
}

func deleteAllPolicies() {
	policies := getPoliciesInRole()

	for _, v := range policies {
		_, err := client.DeleteRolePolicy(context.Background(), &iam.DeleteRolePolicyInput{
			RoleName:   aws.String(readfiles.GetNameFunctionLambda() + "-role"),
			PolicyName: aws.String(v),
		})
		if err != nil {
			panic(err)
		}
	}
}

func (r *RolePolicies) putPolicies() {
	_, err := client.PutRolePolicy(context.Background(), &iam.PutRolePolicyInput{
		RoleName:       aws.String(readfiles.GetNameFunctionLambda() + "-role"),
		PolicyName:     aws.String(r.GetNamePolici()),
		PolicyDocument: aws.String(r.GetBytesFile()),
	})
	if err != nil {
		panic(err)
	}
}

func getPoliciesInRole() []string {
	resp, err := client.ListRolePolicies(context.Background(), &iam.ListRolePoliciesInput{
		RoleName: aws.String(readfiles.GetNameFunctionLambda() + "-role"),
	})

	if err != nil {
		panic(err)
	}

	var policies []string
	policies = append(policies, resp.PolicyNames...)

	return policies
}

func (r *RolePolicies) existPolicieInRol() bool {
	resp := getPoliciesInRole()

	for _, v := range resp {
		if v == r.GetNamePolici() {
			return true
		}
	}
	return false
}

func (r *RolePolicies) deletePolici() {
	_, err := client.DeleteRolePolicy(context.Background(), &iam.DeleteRolePolicyInput{
		PolicyName: aws.String(r.namePolici),
		RoleName:   aws.String(readfiles.GetNameFunctionLambda() + "-role"),
	})
	if err != nil {
		panic(err)
	}
}

func (r *RolePolicies) createPolicie() {
	_, err := client.PutRolePolicy(context.Background(), &iam.PutRolePolicyInput{
		PolicyName:     aws.String(r.namePolici),
		RoleName:       aws.String(readfiles.GetNameFunctionLambda() + "-role"),
		PolicyDocument: aws.String(r.GetBytesFile()),
	})
	if err != nil {
		panic(err)
	}
}
