package readfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type GlobalEnvironment struct {
	PrivateEnvironmentVariables string `json:"privates_environment_variables"` // path of the file
	PathBaseApi                 string `json:"path_base_api"`
}

func newGlobalEnvironment() *GlobalEnvironment {

	// read file json and assign to struct
	var le GlobalEnvironment

	//open  from file
	file, err := os.OpenFile(path.Join(pathGlobalEnv, NameGlobalEnv), os.O_RDONLY, 0644)
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

func (g *GlobalEnvironment) getHiddenEnvironment() map[string]string {
	m := make(map[string]string)

	fmt.Printf("g.PrivateEnvironmentVariables: %v\n", g.PrivateEnvironmentVariables)
	file, err := os.ReadFile(g.PrivateEnvironmentVariables)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func GetPathBase() string {
	return GEnv.PathBaseApi
}
