package proxy

import "fmt"

type Server interface {
	Echo() string
}

type JPServer struct {
	Address string
}

func NewJPServer(address string) *JPServer {
	return &JPServer{
		Address: address,
	}
}

func (s *JPServer) Echo() string {
	return fmt.Sprintf("Echo form: %s", s.Address)
}
