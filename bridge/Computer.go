package bridge

import "fmt"

type PC interface {
	Install(s Soft)
	Run()
}

type ASUS struct {
	Brand string
	App   Soft
}

func (a *ASUS) New() *ASUS {
	a.Brand =  "ASUS"
	return a
}

func (a *ASUS) Install(s Soft) {
	fmt.Printf("%s Mobile Install\n", a.Brand)
	a.App = s
}

func (a *ASUS) Run() {
	a.App.Run()
}

type NOKIA struct {
	Brand string
	App   Soft
}

func (n *NOKIA) New() *NOKIA {
	n.Brand =  "NOKIA"
	return n
}

func (n *NOKIA) Install(s Soft) {
	fmt.Printf("%s Mobile Install\n", n.Brand)
	n.App = s
}

func (n *NOKIA) Run() {
	n.App.Run()
}
