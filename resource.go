package main

import (
	"errors"
	"fmt"
)

func NewResource(vars map[string]string) (Fetcher, error) {
	resourceType := vars["resourceType"]

	switch resourceType {
	case "file":
		return &File{Path: makeUserFilePath(vars["user"], UserFiles[vars["resourceName"]])}, nil
	case "command":
		// ToDO
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("resourceType=%s does not match any defined types", resourceType))
}
