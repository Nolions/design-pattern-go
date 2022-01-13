package flyweight

type Equipment interface {
	GetColor() string
}

type Saber struct {
	color string
}

func (t *Saber) GetColor() string {
	return t.color
}

func NewSaber() *Saber {
	return &Saber{color: "red"}
}

type Archer struct {
	color string
}

func (c *Archer) GetColor() string {
	return c.color
}

func NewArcher() *Archer {
	return &Archer{color: "green"}
}
