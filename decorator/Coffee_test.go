package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMickCoffee(t *testing.T) {
	var c Drink
	c = Coffee{description: "American Coffee"}
	c = Milk{
		description: "Milk",
		drink:       c,
	}

	assert.Equal(t, c.Description(), "American Coffee, Milk")
	assert.Equal(t, c.Cost(), 110)
}

func TestDoubleMickCoffee(t *testing.T) {
	var c Drink
	c = Coffee{description: "American Coffee"}
	c = Milk{
		description: "Milk",
		drink:       c,
	}
	c = Milk{
		description: "Milk",
		drink:       c,
	}

	assert.Equal(t, c.Description(), "American Coffee, Milk, Milk")
	assert.Equal(t, c.Cost(), 120)
}

func TestBlackTeaMickCoffee(t *testing.T) {
	var c Drink
	c = Coffee{description: "American Coffee"}
	c = Milk{
		description: "Milk",
		drink:       c,
	}
	c = BlackTea{
		description: "Black Tea",
		drink:       c,
	}

	assert.Equal(t, c.Description(), "American Coffee, Milk, Black Tea")
	assert.Equal(t, c.Cost(), 115)
}
