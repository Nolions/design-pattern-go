package chainOfResponsibility

import "fmt"

type GeneralManager struct {
	M manager
}

func (gm *GeneralManager) Allow(r *Request) {
	fmt.Printf("General Manager Allow Request")
	r.GeneralManagerAllow = true
}

func (gm *GeneralManager) SetManager(m manager) {
	gm.M = m
}
