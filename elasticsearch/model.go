package elasticsearch

type Model interface {
	GetId() uint64
	Name() string
}
