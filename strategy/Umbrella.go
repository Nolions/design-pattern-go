package strategy


type Umbrella interface {
	price(count int) float32
}

type Rain struct {
	cost float32
}

func (u *Rain) price(count int) float32 {
	return u.cost * float32(count) * 1
}


type Sun struct {
	cost float32
}

func (u *Sun) price(count int) float32 {
	return u.cost * float32(count) * 0.5
}