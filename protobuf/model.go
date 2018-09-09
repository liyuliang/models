package protobuf

type Model interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)
}
