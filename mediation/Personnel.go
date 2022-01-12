package mediation

import "fmt"

type Personnel interface {
	Say(msg string)
	Notify(msg string)
}

type PM struct {
	name string
	mediator Mediator
}

func (c *PM) New(m Mediator) *PM {
	c.name = "A"
	c.mediator = m

	return c
}

func (c *PM) Say(msg string) {
	fmt.Printf("%s say %s\n",c.name,  msg)
	c.mediator.Say(c, msg)
}

func (c *PM) Notify(msg string) {
	fmt.Printf("%s get %s\n",c.name,  msg)
}

type RD struct {
	name string
	mediator Mediator
}

func (c *RD) New(m Mediator) *RD {
	c.name = "B"
	c.mediator = m

	return c
}

func (c *RD) Say(msg string) {
	fmt.Printf("%s say %s\n",c.name,  msg)
	c.mediator.Say(c, msg)
}

func (c *RD) Notify(msg string) {
	fmt.Printf("%s get %s\n",c.name,  msg)
}
