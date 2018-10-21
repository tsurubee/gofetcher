package model

import (
	"errors"
	"fmt"
)

func NewResource(vars map[string]string) (Fetcher, error) {
	resourceType := vars["resourceType"]

	switch resourceType {
	case "file":
		path, err := makeUserFilePath(vars["user"], UserFiles[vars["resourceName"]])
		if err != nil {
			return nil, err
		}
		return &File{Path: path}, nil
	case "command":
		// ToDO
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("resourceType=%s does not match any defined types", resourceType))
}
