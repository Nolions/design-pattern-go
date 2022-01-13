package flyweight

type Game struct {
	saber  []*Role
	archer []*Role
}

func NewGame() *Game {
	return &Game{
		saber:  make([]*Role, 1),
		archer: make([]*Role, 1),
	}
}

func (c *Game) AddSaber() {
	player := NewRole("T", SaberType)
	c.saber = append(c.saber, player)
	return
}

func (c *Game) AddArcher() {
	player := NewRole("CT", ArcherType)
	c.archer = append(c.archer, player)
	return
}
