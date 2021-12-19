package builder

import "image/color"

type Type int

const (
	Auto Type = iota
	Motorcycle
)

type Transportation interface {
	SetType(t Type) Transportation
	SetColor(color color.Color) Transportation
	SetBrand(brand string) Transportation
	WheelCount() Transportation
	//Build()
}
