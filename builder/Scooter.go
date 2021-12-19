package builder

import "image/color"

type Scooter struct {
	Brand string
	Color color.Color
	Type  Type
	Wheel int
}

func NewScooter() *Scooter {
	return &Scooter{}
}

func (s *Scooter) SetType(t Type) Transportation {
	s.Type = t
	return s
}

func (s *Scooter) SetColor(color color.Color) Transportation {
	s.Color = color
	return s
}

func (s *Scooter) SetBrand(brand string) Transportation {
	s.Brand = brand
	return s
}

func (s *Scooter) WheelCount() Transportation {
	s.Wheel = 2
	return s
}
