package simpleFactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIphone(t *testing.T) {
	p := NewPhone("iphone")

	assert.Equal(t, p.OS(), "IOS")
	assert.Equal(t, p.Price(), 30000)
}


func TestAndroid(t *testing.T) {
	p := NewPhone("android")

	assert.Equal(t, p.OS(), "Android")
	assert.Equal(t, p.Price(), 20000)
}