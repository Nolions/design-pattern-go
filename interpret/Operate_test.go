package interpret

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestANDOperate(t *testing.T) {
	a := NewOperate([]string{"a", "b"})
	b := NewOperate([]string{"c", "d"})

	and := NewOrOperate(a, b)

	assert.Equal(t, and.Interpret("a"), true)
	assert.Equal(t, and.Interpret("b"), true)
}

func TestOROperate(t *testing.T) {
	a := NewOperate([]string{"a", "b"})
	b := NewOperate([]string{"c", "a"})

	and := NewAndOperate(a, b)

	assert.Equal(t, and.Interpret("a"), true)
}
