package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJan(t *testing.T) {
	m := newMonth("Jan", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestFeb(t *testing.T) {
	m := newMonth("Feb", 28)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), true)
	assert.Equal(t, m.NonLeap(), false)
}

func TestLeapFeb(t *testing.T) {
	m := newMonth("Feb", 29)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), true)
}

func TestMar(t *testing.T) {
	m := newMonth("Mar", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestApr(t *testing.T) {
	m := newMonth("Apr", 30)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), true)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestMay(t *testing.T) {
	m := newMonth("Apr", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestJune(t *testing.T) {
	m := newMonth("Apr", 30)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), true)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestJuly(t *testing.T) {
	m := newMonth("Apr", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestAug(t *testing.T) {
	m := newMonth("Apr", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestSept(t *testing.T) {
	m := newMonth("Apr", 30)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), true)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestOct(t *testing.T) {
	m := newMonth("Apr", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestNov(t *testing.T) {
	m := newMonth("Nov", 30)

	assert.Equal(t, m.Solar(), false)
	assert.Equal(t, m.Lunar(), true)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}

func TestDec(t *testing.T) {
	m := newMonth("Apr", 31)

	assert.Equal(t, m.Solar(), true)
	assert.Equal(t, m.Lunar(), false)
	assert.Equal(t, m.Leap(), false)
	assert.Equal(t, m.NonLeap(), false)
}