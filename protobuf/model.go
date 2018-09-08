package protobuf

type Model interface {
	GetId() uint64
	Reset()
	String() string
	ProtoMessage()
}
