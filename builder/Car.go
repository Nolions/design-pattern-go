package builder

import "image/color"

type Car struct {
	Brand string
	Color color.Color
	Type  Type
	Wheel int
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) SetType(t Type) Transportation {
	c.Type = t
	return c
}

func (c *Car) SetColor(color color.Color) Transportation {
	c.Color = color
	return c
}

func (c *Car) SetBrand(brand string) Transportation {
	c.Brand = brand
	return c
}

func (c *Car) WheelCount() Transportation {
	c.Wheel = 4
	return c
}

//func (c *Car) Build() {
//
//}
