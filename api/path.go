package api

import (
	"context"
	readfiles "deploye-lambda/read_files"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

type PathApi struct {
	completePath string
	id           string
}

// constructo pathapi
func NewPathApi() *PathApi {
	// find count of subdirectories is the base path base in the file env.json
	count := readfiles.CountManySubdirectoriesExistFile(readfiles.NameGlobalEnv)

	// get the current path
	path, _ := os.Getwd()

	// splith path from end to start with the count
	pathSplit := strings.Join(strings.Split(path, "/")[len(strings.Split(path, "/"))-count:len(strings.Split(path, "/"))-1], "/")
	fmt.Printf("pathSplit: %v\n", pathSplit)

	return &PathApi{
		completePath: pathSplit,
	}
}

// get id from pathapi, if not exist then consulte to create resources.
// if no create resoources, then end program to deploy api, but is need this for deploy method and other thins
func (p *PathApi) SetId() {
	if p.id == "" {
		p.id = getApiGateway("contabilidad/" + p.completePath)
		return
	}
	return
}

func getApiGateway(path string) string {

	// find id from path
	res, err := client.GetResources(context.Background(), &apigateway.GetResourcesInput{
		RestApiId: aws.String(readfiles.ApiId),
	})

	if err != nil {
		panic(err)
	}

	pathSplited := strings.Split(path, "/")

	existPath := func(subPath string) (string, bool) {
		for _, r := range res.Items {
			if ok, count := subStringFromInit(subPath, *r.Path); ok {
				switch count {
				case 0:
					return *r.Id, true
				case 1:
					return *r.ParentId, true
				}
			}
		}

		return "", false
	}

	// range pathSplited
	acumPath := ""
	id := ""
	for _, val := range pathSplited {
		acumPath += "/" + val
		if rid, ok := existPath(acumPath); ok {
			fmt.Println("exist path " + acumPath)
			id = rid
		} else {
			// consulte in console with color blue if want deploy resource
			fmt.Printf("\033[1;34m%s\033[0m\n", "do you want deploy resource "+acumPath+"? (Y/n)")
			var answer string
			fmt.Scanln(&answer)
			if answer == "" || answer == "y" {
				fmt.Println("deploy resource " + acumPath)
				if id, err = createResource(val, id); err != nil {
					panic(err)
				}
			} else {
				fmt.Println("end program")
				os.Exit(0)
			}
		}
	}

	return id
}

// create resources and return the id of created resource and error
func createResource(path string, parentId string) (string, error) {
	// create resources
	resp, err := client.CreateResource(context.Background(), &apigateway.CreateResourceInput{
		PathPart:  aws.String(path),
		RestApiId: aws.String(readfiles.ApiId),
		ParentId:  aws.String(parentId),
	})

	if err != nil {
		return "", err
	}
	return *resp.Id, nil
}

// compare string since init to end children with parent string, if is always equal return true
// if exist, then, find count time parent have
func subStringFromInit(children, parent string) (bool, int) {
	// string to char
	initChar := []rune(children)
	endChar := []rune(parent)

	// range initChar
	for i, val := range initChar {
		if i < len(endChar) {
			if val != endChar[i] {
				return false, 0
			}
		} else {
			return false, 0
		}
	}

	replaced := strings.ReplaceAll(parent, children, "")
	getCount := func() int {
		count := 0
		for _, v := range strings.Split(replaced, "/") {
			if v != "" {
				count++
			}
		}
		return count
	}
	fmt.Printf("parent: %v\n", parent)
	return true, getCount()
}

// get method for the current resource
// need have the id of resource previous to call this function
func (p *PathApi) getMethod() map[string]types.Method {
	if p.id != "" {
		// get method from resource
		res, err := client.GetResource(context.Background(), &apigateway.GetResourceInput{
			ResourceId: aws.String(p.id),
			RestApiId:  aws.String(readfiles.ApiId),
		})

		if err != nil {
			panic(err)
		}
		return res.ResourceMethods
	} else {
		panic("no exist id for resource, try call this function after get id")
	}
}
