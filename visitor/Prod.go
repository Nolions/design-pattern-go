package visitor

import "fmt"

type Prod struct {

}

func (v Prod) Visit() string {
	fmt.Println("Env Prod")
	return "prod"
}
