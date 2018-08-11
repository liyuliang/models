package database

type modelCreator func() model

var creatorList []modelCreator

func Register(method modelCreator) {

	creatorList = append(creatorList, method)
}

type modelList map[string]modelCreator

var _list modelList

func List() modelList {

	if len(_list) != len(creatorList) {

		_list = make(map[string]modelCreator)

		for _, creator := range creatorList {
			m := creator()
			_list[m.Name()] = creator
		}
	}

	return _list
}
