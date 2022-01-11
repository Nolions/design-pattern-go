package chainOfResponsibility

import "fmt"

type Director struct {
	M manager
}

func (d *Director) Allow(r *Request) {
	if r.Count <5 {
		fmt.Printf("Director Allow Request")
		fmt.Printf("Your Reques is Allow")
		r.DirectorAllow = true
	} else {
		d.M.Allow(r)
	}
}

func (d *Director) SetManager(m manager) {
	d.M = m
}
