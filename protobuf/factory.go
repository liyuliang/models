package protobuf

import (
	"strings"
	"errors"
)

func Get(modelName string) (Model, error) {

	l := List()

	for name, creator := range l {

		if strings.ToLower(name) == strings.ToLower(modelName) {
			m := creator()
			return m, nil
		}
	}
	return nil, errors.New("can not get the Model by name :" + modelName)
}
