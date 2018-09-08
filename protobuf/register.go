package protobuf

func init() {
	_list = make(map[string]modelCreator)

	_list["manhua"] = func() Model {
		return new(Manhua)
	}
}
