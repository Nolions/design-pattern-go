package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRain(t *testing.T)  {
	w := Weather{}
	w.New(RainEnum)
	price := w.Sell(1)

	assert.Equal(t, price, float32(10))
}

func TestSun(t *testing.T)  {
	w := Weather{}
	w.New(SunEnum)
	price := w.Sell(1)

	assert.Equal(t, price, float32(5))
}