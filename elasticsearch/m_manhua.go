package elasticsearch

import "time"

func init() {
	Register(func() Model {
		return new(Manhua)
	})
}

type Manhua struct {
	Id      uint
	Site    string
	Number  int
	Chapter string
	Title   string
	Update  time.Time
}

func (table *Manhua) Name() string {
	return "MANHUA"
}
