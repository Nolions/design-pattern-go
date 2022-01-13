package flyweight

import "fmt"

const (
	SaberType  = "saber"
	ArcherType = "archer"
)

var (
	equipment = &EquipmentFactory{
		EquipmentMap: make(map[string]Equipment),
	}
)

type EquipmentFactory struct {
	EquipmentMap map[string]Equipment
}

func (d *EquipmentFactory) GetEquipment(dressType string) (Equipment, error) {
	if d.EquipmentMap[dressType] != nil {
		return d.EquipmentMap[dressType], nil
	}

	if dressType == SaberType {
		d.EquipmentMap[dressType] = NewSaber()
		return d.EquipmentMap[dressType], nil
	}
	if dressType == ArcherType {
		d.EquipmentMap[dressType] = NewArcher()
		return d.EquipmentMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetEquipmentFactory() *EquipmentFactory {
	return equipment
}