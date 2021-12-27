package template

import "fmt"

type AgeOfEmpires struct {
}

func (AgeOfEmpires) initialize() {
	fmt.Println("initialize")
}

func (AgeOfEmpires) start() {
	fmt.Println("start")
}

func (AgeOfEmpires) end() {
	fmt.Println("end")
}
