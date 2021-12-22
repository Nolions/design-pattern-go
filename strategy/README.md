# Strategy 策略模式

## 策略決定

在策略模式中有兩種使用方式

1. 用戶直接決定使策略，也是範例中所使用的方式，這種方式可以去除掉Strategy中`if`或`switch`運用，讓Strategy中的看起來比較優雅。

```go
// strategy
func (w *Weather) New(u Umbrella) {
	w.u = u
}

// client
w := Weather{}
w.New(RainEnum)
price := w.Sell(1)

w := Weather{}
w.New(RainEnum)
price := w.Sell(1)
```

2. 用戶不直接傳入策略，而是類似傳入enum方式讓Strategy去決定使用什麼策略，也是範例中所使用的方式，在某些策略決定Enum可能是全域變數或是其他不用透過傳入值情況下很適合使用。

```go
// strategy
func (w *Weather) New(e WeatherEnum) {
	switch e {
	case SunEnum:
		w.u = &Sun{cost: 10}
	case RainEnum:
		w.u = &Rain{cost: 10}

	}
}

// client
w := Weather{}
w.New(RainEnum)
price := w.Sell(1)

w := Weather{}
w.New(RainEnum)
price := w.Sell(1)
```