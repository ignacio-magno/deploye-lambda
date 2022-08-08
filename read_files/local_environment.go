package readfiles

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LocalEnvironment struct {
	LambdaFunctionName   string `json:"lambda_function_name"`
	EnvironmentVariables []struct {
		Name string `json:"name"`
	} `json:"environment_variables"`
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

func GetNameFunctionLambda() string {
	return LEnv.LambdaFunctionName
}
