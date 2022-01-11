package chainOfResponsibility

type Request struct {
	Count               int
	DirectorAllow       bool
	GeneralManagerAllow bool
	ManagerAllow        bool
}

func NewRequest(c int) *Request {
	return &Request{
		Count:               c,
		DirectorAllow:       false,
		ManagerAllow:        false,
		GeneralManagerAllow: false,
	}
}

type manager interface {
	Allow(*Request)
	SetManager(manager)
}
