package protobuf

import "models/protobuf/task"

func init() {
	_list = make(map[string]modelCreator)

	_list["manhua"] = func() Model {
		return new(Manhua)
	}
	_list["a"] = func() Model {
		return new(task.SaverManhuaSite)
	}
}
