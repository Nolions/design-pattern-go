package flyweight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	g := NewGame()
	g.AddSaber()
	g.AddSaber()

	g.AddArcher()
	g.AddArcher()
	g.AddArcher()

	d := GetEquipmentFactory()

	for eType, e := range d.EquipmentMap {
		fmt.Printf("Role: %s\nColor: %s\n", eType, e.GetColor())
	}

	assert.Equal(t, len(d.EquipmentMap), 2)
}
