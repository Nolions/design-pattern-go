package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxyServer(t *testing.T) {
	p := NewProxyServer("localhost")
	p.NewJPServer()

	ping := p.Echo()

	assert.Equal(t, ping, "Echo form: localhost")
}
