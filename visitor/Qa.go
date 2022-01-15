package visitor

import "fmt"

type Qa struct {

}

func (q Qa) Visit() string {
	fmt.Println("Env Qa")
	return "qa"
}