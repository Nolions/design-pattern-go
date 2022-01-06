package memento

type Role struct {
	state string
}

func (r *Role) Save() *Memento {
	return &Memento{state: r.state}
}

func (r *Role) Load(m *Memento) {
	r.state = m.GetSavedState()
}

func (r *Role) setState(state string) {
	r.state = state
}

func (r *Role) getState() string {
	return r.state
}