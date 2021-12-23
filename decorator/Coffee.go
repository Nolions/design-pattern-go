package decorator

import "fmt"

type Drink interface {
	Cost() int
	Description() string
}

type Coffee struct {
	description string
}

func (c Coffee) Description() string {
	return c.description
}

func (c Coffee) Cost() int {
	return 100
}

type Milk struct {
	drink       Drink
	description string
}

func (m Milk) Description() string {
	return fmt.Sprintf("%s, %s", m.drink.Description(), m.description)
}

func (m Milk) Cost() int {
	return m.drink.Cost() + 10
}

type BlackTea struct {
	drink       Drink
	description string
}

func (b BlackTea) Description() string {
	return fmt.Sprintf("%s, %s", b.drink.Description(), b.description)
}

func (b BlackTea) Cost() int {
	return b.drink.Cost() + 5
}