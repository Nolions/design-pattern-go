package mediation

import "testing"

func TestMediator(t *testing.T)  {
	m := NewUIMediator()

	a := &PM{}
	a.New(m)

	b := &RD{}
	b.New(m)

	m.a = a
	m.b = b

	a.Say("123")
}
