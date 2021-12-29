package observer

import "fmt"

type Customer interface {
	update()
}

type CustomerA struct {
	name string
}

func (c *CustomerA) update() {
	fmt.Printf("Customer %s get notification \n", c.name)
}

type CustomerB struct {
	name string
}

func (c *CustomerB) update() {
	fmt.Printf("Customer %s get notification \n", c.name)
}
