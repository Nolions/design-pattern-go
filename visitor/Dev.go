package visitor

import "fmt"

type Dev struct {

}

func (d Dev) Visit() string {
	fmt.Println("Env Dev")
	return "dev"
}
