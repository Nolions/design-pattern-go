package chainOfResponsibility

import "fmt"

type Manager struct {
	M manager
}

func (m *Manager) Allow(r *Request) {
	if r.Count <10 {
		fmt.Printf("Manager Allow Request")
		r.ManagerAllow = true
	} else {
		m.M.Allow(r)
	}
}

func (mer *Manager) SetManager(m manager) {
	mer.M = m
}
