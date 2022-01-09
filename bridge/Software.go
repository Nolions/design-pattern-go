package bridge

import "fmt"

type Soft interface {
	Run()
}

type Game struct {
	Type string
}

func (g *Game) New() *Game {
	g.Type = "Game"
	return g
}

func (g *Game) Run() {
	fmt.Printf("Run %s Application\n", g.Type)
}

type Directory struct {
	Type string
}

func (d *Directory) New() *Directory {
	d.Type = "Directory"
	return d
}

func (d *Directory) Run() {
	fmt.Printf("Run %s Application\n", d.Type)
}
