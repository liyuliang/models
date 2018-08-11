package database

import (
	"testing"
	"gitee.com/liyuliang/utils/access"
	"encoding/json"
)

func TestList(t *testing.T) {

	l := List()

	if len(l) == 0 {
		t.Error("Model list should not be empty")
	}
}

type model1 struct {
	Author string
}
type model2 struct {
	Author string
}
type model3 struct {
	Author string
}

func (model3 *model3) Name() string {
	return "model3"
}
func (model2 *model2) Name() string {
	return "model2"
}
func (model1 *model1) Name() string {
	return "model1"
}

func TestGet(t *testing.T) {
	Register(func() model {
		return new(model1)
	})
	Register(func() model {
		return new(model2)
	})
	Register(func() model {
		return new(model3)
	})

	modelName := "model1"
	m, err := Get(modelName)
	if err != nil {
		t.Error(err.Error())
	}
	if m.Name() != "model1" {
		t.Error("Get model1 failed, get the " + m.Name() + " Model")
	}

	modelName = "model2"
	m, err = Get(modelName)
	if err != nil {
		t.Error(err.Error())
	}
	if m.Name() != modelName {
		t.Error("Get " + modelName + " failed, get the " + m.Name() + " Model")
	}

	modelName = "model3"
	m, err = Get(modelName)

	if err != nil {
		t.Error(err.Error())
	}
	if m.Name() != modelName {
		t.Error("Get " + modelName + " failed, get the " + m.Name() + " Model")
	}
}

func TestGetNotAPointer(t *testing.T) {

	Register(func() model {
		return new(model1)
	})

	data := make(map[string]interface{})
	data["Author"] = "liang"

	modelName := "model1"
	m, _ := Get(modelName)
	m2, _ := Get(modelName)

	access.Set(m, data)
	mJson, _ := json.Marshal(m)
	m2Json, _ := json.Marshal(m2)

	if string(mJson) == string(m2Json) {
		t.Error("the factory get method return a pointer, m1 has value, but m2 has not")
	}

	m3, _ := Get(modelName)
	m3Json, _ := json.Marshal(m3)

	if string(mJson) == string(m3Json) {
		t.Error("the factory get method return a pointer, m1 has value, but m3 has not")
	}

	if string(m2Json) != string(m3Json) {
		t.Error("the factory get method return a pointer, m2 m3 should has not value")
	}

}
