package command

type Button struct {
	Command Command
}

func (b *Button) press() {
	b.Command.Execute()
}

type Device interface {
	On()
	Off()
}

type Control struct {
	PowerOn bool
}

func NewControl() *Control {
	return &Control{
		PowerOn: false,
	}
}

func (c *Control) On() {
	c.PowerOn = true
}

func (c *Control) Off() {
	c.PowerOn = false
}
