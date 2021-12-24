package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTCTaiwanPhone(t *testing.T) {
	f := CreateFactory(HTCTaiwanPhoneEnum)
	m :=f.Create().Market()

	assert.Equal(t, m, "Your HTC Phone made in Taiwan")
}

func TestHTCChinaPhone(t *testing.T) {
	f := CreateFactory(HTCChinaPhoneEnum)
	m :=f.Create().Market()

	assert.Equal(t, m, "Your HTC Phone made in China")
}

func TestASUSTaiwanPhone(t *testing.T) {
	f := CreateFactory(ASUSTaiwanPhoneEnum)
	m :=f.Create().Market()

	assert.Equal(t, m, "Your ASUS Phone made in Taiwan")
}

func TestASUSChinaPhone(t *testing.T) {
	f := CreateFactory(ASUSChinaPhoneEnum)
	m :=f.Create().Market()

	assert.Equal(t, m, "Your ASUS Phone made in China")
}