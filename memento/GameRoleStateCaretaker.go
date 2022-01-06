package memento

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) Save(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) Load(index int) *Memento {
	return c.mementos[index]
}