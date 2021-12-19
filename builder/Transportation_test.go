package builder

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestCar(t *testing.T) {
	c := NewCar()
	c.SetBrand("toyota").SetType(Auto).SetColor(color.White).WheelCount()

	assert.Equal(t, c.Brand, "toyota")
	assert.Equal(t, c.Type, Auto)
	assert.Equal(t, c.Color, color.White)
	assert.Equal(t, c.Wheel, 4)
}

func TestScooter(t *testing.T) {
	s := NewScooter()
	s.SetBrand("yamaha").SetType(Motorcycle).SetColor(color.Black).WheelCount()

	assert.Equal(t, s.Brand, "yamaha")
	assert.Equal(t, s.Type, Motorcycle)
	assert.Equal(t, s.Color, color.Black)
	assert.Equal(t, s.Wheel, 2)
}
