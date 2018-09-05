package protobuf

type modelCreator func() model

var creatorList []modelCreator

func Register(method modelCreator) {

	creatorList = append(creatorList, method)
}

type modelList map[string]modelCreator

var _list modelList

func List() modelList {
	return _list
}

func init() {
	_list = make(map[string]modelCreator)

	_list["manhua"] = func() model {
		return new(Manhua)
	}
}
