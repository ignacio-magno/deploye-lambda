package readfiles

import (
	"fmt"
	"os"
	"os/exec"
)

// build the code go and comprime it
func DeployCode() ([]byte, error) {
	// elimina el previo codigo o zip si es que existe
	deleteCode()
	deleteZip()

	// construye el codigo
	buildCode()
	buildZip()

	// open file zip
	f, err := os.ReadFile(handlerZip)
	if err != nil {
		return nil, err
	}

	// get bytes of file.reader

	return f, nil
}

func buildCode() {
	// execut shell command
	// go build -o /tmp/code.zip -v
	err := exec.Command("go", "build", "main.go").Run()
	if err != nil {
		panic(err)
	}
}

func deleteCode() {
	err := exec.Command("rm", "main").Run()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
}

func buildZip() {
	err := exec.Command("zip", "main.zip", "main").Run()
	if err != nil {
		panic(err)
	}
}

func deleteZip() {
	err := exec.Command("rm", "main.zip").Run()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
}
