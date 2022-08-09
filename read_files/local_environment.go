package readfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LocalEnvironment struct {
	LambdaFunctionName   string            `json:"lambda_function_name"`
	EnvironmentVariables map[string]string `json:"environment_variables"`
}

func newLocalEnvironment() *LocalEnvironment {

	// read file json and assign to struct
	var le LocalEnvironment

	//open  from file
	file, err := os.OpenFile(pathLocalEnv+NameLocalEnv, os.O_RDONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	// close file
	defer file.Close()

	// get bytes from file
	bytes, err := ioutil.ReadAll(file)
	// check error
	if err != nil {
		panic(err)
	}

	// marsahl file to struct
	err = json.Unmarshal(bytes, &le)
	if err != nil {
		panic(err)
	}
	return &le
}

func GetEnvironmentVariables() map[string]string {
	toReturn := make(map[string]string)
	// credentials
	// consulte if want to use the environment variables privates
	fmt.Println("you wish use the privates environment variables y/N")
	var answer string
	fmt.Scanln(&answer)
	if answer == "y" {
		m := GEnv.getHiddenEnvironment()
		for k, v := range m {
			toReturn[k] = v
		}
	}

	for k, v := range LEnv.EnvironmentVariables {
		toReturn[k] = v
	}

	return toReturn
}

func GetNameFunctionLambda() string {
	return LEnv.LambdaFunctionName
}
