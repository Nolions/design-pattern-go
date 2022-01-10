package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestControl(t *testing.T) {
	c := NewControl()

	onCMD := &OnCommend{
		Device: c,
	}

	offCMD := &OffCommend{
		Device: c,
	}

	onBtn := &Button{
		Command: onCMD,
	}


	offBtn := &Button{
		Command: offCMD,
	}

	assert.Equal(t, c.PowerOn, false)

	press(onBtn)
	assert.Equal(t, c.PowerOn, true)

	press(offBtn)
	assert.Equal(t, c.PowerOn, false)
}

func press(btn *Button) {
	btn.press()
}
