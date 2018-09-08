package elasticsearch

func init() {
	Register(func() Model {
		return new(Manhua)
	})
}

type Manhua struct {
	Id      uint64
	Site    string
	Number  int32
	Chapter string
	Title   string
	Update  string
}

func (m *Manhua) Name() string {
	return "MANHUA"
}

func (m *Manhua) GetId() uint64 {
	return m.Id
}
