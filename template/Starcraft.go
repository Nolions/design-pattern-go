package template

import "fmt"

type Starcraft struct {

}

func (Starcraft) initialize() {
	fmt.Println("initialize")
}

func (Starcraft) start() {
	fmt.Println("start")
}

func (Starcraft) end() {
	fmt.Println("end")
}
