package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var c *Caretaker

func init() {
	c = &Caretaker{
		mementos: make([]*Memento, 0),
	}
}

func TestGameRole(t *testing.T) {
	r := &Role{
		state: "A",
	}
	c.Save(r.Save())

	r.setState("B")
	c.Save(r.Save())

	r.setState("C")
	c.Save(r.Save())

	r.Load(c.Load(0))
	assert.Equal(t, r.getState(), "A")

	r.Load(c.Load(1))
	assert.Equal(t, r.getState(), "B")

	r.Load(c.Load(2))
	assert.Equal(t, r.getState(), "C")
}
