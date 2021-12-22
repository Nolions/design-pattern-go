package strategy

type Weather struct {
	u Umbrella
}

type WeatherEnum int

const (
	SunEnum = iota
	RainEnum
)

func (w *Weather) New(e WeatherEnum) {
	switch e {
	case SunEnum:
		w.u = &Sun{cost: 10}
	case RainEnum:
		w.u = &Rain{cost: 10}

	}
}

//func (w *Weather) New(u Umbrella) {
//	w.u = u
//}

func (w *Weather) Sell(count int) float32 {
	return w.u.price(count)
}
