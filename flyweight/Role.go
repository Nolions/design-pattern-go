package flyweight

type Role struct {
	equipment Equipment
	roleType  string
}

func NewRole(rType, eType string) *Role {
	e, _ := GetEquipmentFactory().GetEquipment(eType)
	return &Role{
		roleType:  rType,
		equipment: e,
	}
}
