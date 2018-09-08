package elasticsearch

import "time"

func init() {
	Register(func() Model {
		return new(Manhua)
	})
}

type Manhua struct {
	Id      uint64
	Site    string
	Number  int
	Chapter string
	Title   string
	Update  time.Time
}

func (m *Manhua) Name() string {
	return "MANHUA"
}

func (m *Manhua) GetId() uint64 {
	return m.Id
}
