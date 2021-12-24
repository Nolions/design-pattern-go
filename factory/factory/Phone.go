package factory

import "fmt"

type Phone interface {
	Market() string
}

type HTC struct {
	Origin string
}

func (h HTC) Market() string {
	return fmt.Sprintf("Your HTC Phone made in %s", h.Origin)
}

type ASUS struct {
	Origin string
}

func (a ASUS) Market() string  {
	return fmt.Sprintf("Your ASUS Phone made in %s", a.Origin)
}