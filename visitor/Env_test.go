package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProdEnv(t *testing.T) {
	env := Prod{}
	assertEqual(t, env)
}

func TestDevEnv(t *testing.T) {
	env := Dev{}
	assertEqual(t, env)
}

func TestQaEnv(t *testing.T) {
	env := Qa{}
	assertEqual(t, env)
}

func assertEqual(t *testing.T, env Visitor) {
	e := Element{}
	a := e.accept(env)

	m := Env{}
	p := m.Print(env)

	assert.Equal(t, a, p)
}
