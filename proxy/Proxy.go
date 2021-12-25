package proxy

type Proxy struct {
	Server  *JPServer
	Address string
}

func NewProxyServer(address string) *Proxy {
	return &Proxy{
		Address: address,
	}
}

func (p *Proxy) NewJPServer() {
	if p.Server == nil {
		p.Server = NewJPServer(p.Address)
	}
}

func (p *Proxy) Echo() string {
	return p.Server.Echo()
}


