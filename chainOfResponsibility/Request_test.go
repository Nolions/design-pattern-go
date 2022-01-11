package chainOfResponsibility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var d *Director
var m *Manager
var gm *GeneralManager

func init() {
	d = &Director{}
	m = &Manager{}
	gm = &GeneralManager{}

	m.SetManager(gm)
	d.SetManager(m)
}

func TestDirectorAllowRequest(t *testing.T) {
	r := NewRequest(4)
	d.Allow(r)

	assert.Equal(t, r.DirectorAllow, true)
	assert.Equal(t, r.ManagerAllow, false)
	assert.Equal(t, r.GeneralManagerAllow, false)
}

func TestManagerAllowRequest(t *testing.T) {
	r := NewRequest(9)
	d.Allow(r)

	assert.Equal(t, r.DirectorAllow, false)
	assert.Equal(t, r.ManagerAllow, true)
	assert.Equal(t, r.GeneralManagerAllow, false)
}

func TestGeneralManagerAllowRequest(t *testing.T) {
	r := NewRequest(10)
	d.Allow(r)

	assert.Equal(t, r.DirectorAllow, false)
	assert.Equal(t, r.ManagerAllow, false)
	assert.Equal(t, r.GeneralManagerAllow, true)
}
