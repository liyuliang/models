package protobuf

import (
	"strings"
	"errors"
	"github.com/golang/protobuf/proto"
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

func Marshal(model proto.Message) ([]byte, error) {
	return proto.Marshal(model)
}

func Unmarshal(buf []byte, model proto.Message) error {
	return proto.Unmarshal(buf, model)
}
