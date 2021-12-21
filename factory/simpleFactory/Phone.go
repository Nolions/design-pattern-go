package simpleFactory

type Phone interface {
	OS() string
	Price() int
}

type Iphone struct {
}

func (p *Iphone) OS() string {
	return "IOS"
}

func (p *Iphone) Price() int {
	return 30000
}

type Android struct {

}

func (p *Android) OS() string  {
	return "Android"
}

func (p *Android) Price() int {
	return 20000
}

func NewPhone(p string) Phone {
	if p == "iphone" {
		return &Iphone{}
	} else if p == "android"{
		return &Android{}
	} else {
		return &Android{}
	}
}