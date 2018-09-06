package protobuf

type Model interface {
	Reset()
	String() string
	ProtoMessage()
}
