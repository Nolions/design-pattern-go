package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := NewService()
	s2 := NewService()

	assert.Equal(t, s1, s2)
}
