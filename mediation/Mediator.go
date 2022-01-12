package mediation

type Mediator interface {
	Say(country Personnel, msg string)
}

type UIMediator struct {
	a *PM
	b *RD
}

func NewUIMediator() *UIMediator {
	return &UIMediator{}
}

func (u *UIMediator) Say(country Personnel, msg string) {
	switch country{
	case u.a:
		u.b.Notify(msg)
	case u.b:
		u.b.Notify(msg)
	}
}
