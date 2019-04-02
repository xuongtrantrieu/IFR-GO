package abstract

import "time"

type AbstractModel struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (m *AbstractModel) Clean() {
	m.Created = time.Now()
	m.Modified = time.Now()
}
