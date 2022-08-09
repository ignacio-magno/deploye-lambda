package iamrole

import (
	readfiles "deploye-lambda/read_files"
	"os"
	"path"
	"strconv"
	"strings"
)

type RolePolicies struct {
	pathFile   string
	namePolici string
}

func (r *RolePolicies) GetBytesFile() string {
	f, err := os.ReadFile(r.pathFile)
	if err != nil {
		panic(err)
	}
	return string(f)
}

func (r *RolePolicies) GetNamePolici() string {
	return r.namePolici
}

func (r *RolePolicies) GetPath() string {
	return r.pathFile
}

func NewTrusPolici() *RolePolicies {
	return &RolePolicies{
		pathFile:   readfiles.TrustPoliciPath,
		namePolici: "trust",
	}
}

func NewPutLogPolici() *RolePolicies {
	return &RolePolicies{
		pathFile:   readfiles.PolicyLogsPath,
		namePolici: "logs",
	}
}

func CreateCustomRolePolici(name string, path string) *RolePolicies {
	return &RolePolicies{
		pathFile:   path,
		namePolici: name,
	}
}

func getRolesInFolder() ([]string, func(s string) *RolePolicies) {
	// read dir and get all files
	files, err := os.ReadDir(readfiles.FolderRoles)
	if err != nil {
		panic(err)
	}

	var s []string
	for _, de := range files {
		s = append(s, de.Name())
	}

	return s, func(indiceFile string) *RolePolicies {
		ind, err := strconv.Atoi(indiceFile)
		if err != nil {
			panic(err)
		}

		file := files[ind-1]
		pathFile := path.Join(readfiles.FolderRoles, file.Name())

		return CreateCustomRolePolici(strings.Split(file.Name(), ".")[0], pathFile)
	}
}
