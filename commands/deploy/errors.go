package deploy

import "errors"

var ContainerNotFound = errors.New("container not found")

func IsContainerNotFound(err error) bool {
	return err == ContainerNotFound
}
