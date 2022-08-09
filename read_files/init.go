package readfiles

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

// this package read all files need to obtain data

var (
	pathLocalEnv  = "./"
	pathGlobalEnv = ""
	ApiId         = "m0gqzb21qk"
	LEnv          *LocalEnvironment
	Method        = ""

	handlerMain = "main.go"
	handlerZip  = "main.zip"

	FolderRoles     = "roles"
	TrustPoliciPath = path.Join(FolderRoles, "trust-policy.json")
	PolicyLogsPath  = path.Join(FolderRoles, "policy-logs.json")
)

const (
	NameGlobalEnv = "env.json"
	NameLocalEnv  = "environment.json"
)

func init() {
	// find count of subdirectories is the base path base in the file env.json
	count := CountManySubdirectoriesExistFile(NameGlobalEnv)

	// get the current path
	path, _ := os.Getwd()

	pathGlobalEnv = strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-count], "/")

	// split path with / and get the last element
	Method = strings.ToUpper(strings.Split(path, "/")[len(strings.Split(path, "/"))-1])

	fmt.Printf("Method: %v\n", Method)

	LEnv = newLocalEnvironment()

}

func PrintPath() {
	fmt.Printf("pathGlobalEnv: %v\n", pathGlobalEnv)
}

func CountManySubdirectoriesExistFile(pathFile string) int {

	count := 0
	subDirectories := ""

	for {
		if !existFile(subDirectories + pathFile) {
			subDirectories += "../"
			count++
		} else {
			break
		}
		// time.sleep(1)
		time.Sleep(1 * time.Second)
	}

	return count
}

func existFile(pathFile string) bool {
	// find if exist file
	_, err := os.Stat(pathFile)
	return err == nil
}
