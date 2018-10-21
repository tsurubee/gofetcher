package model

import (
	"errors"
	"fmt"
	"github.com/tsurubee/gofetcher/config"
)

func NewResource(vars map[string]string, c *config.Config) (Fetcher, error) {
	resourceType := vars["resourceType"]

	switch resourceType {
	case "file":
		path, err := makeUserFilePath(vars["user"], c.UserFiles[vars["resourceName"]])
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
