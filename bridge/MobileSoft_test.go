package bridge

import (
	"testing"
)

var sg Game
var sd Directory

var mn NOKIA
var ma ASUS

func init() {
	sg = Game{}
	sd = Directory{}

	mn = NOKIA{}
	ma = ASUS{}
}

func TestNokiaGame(t *testing.T) {
	sg.New()

	mn.New()
	mn.Install(&sg)
	mn.Run()
}

func TestNokiaDirectory(t *testing.T) {
	sd.New()

	mn.New()
	mn.Install(&sd)
	mn.Run()
}

func TestASUSGame(t *testing.T) {
	sg.New()

	ma.New()
	ma.Install(&sg)
	ma.Run()
}

func TestASUSDirectory(t *testing.T) {
	sd.New()

	ma.New()
	ma.Install(&sd)
	ma.Run()
}
