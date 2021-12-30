package abstractFactory

import "testing"

func TestAdidasShoeFactory(t *testing.T) {
	adidasFactory, _ := getSportsFactory("adidas")
	shoe :=adidasFactory.makeShirt()

	shoe.getLogo()
}

func TestAdidasShirtFactory(t *testing.T) {
	adidasFactory, _ := getSportsFactory("adidas")
	shirt :=adidasFactory.makeShirt()

	shirt.getLogo()
}
