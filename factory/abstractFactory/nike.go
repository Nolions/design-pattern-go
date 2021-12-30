package abstractFactory

type nike struct {
}

type nikeShoe struct {
	shoe
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

type nikeShirt struct {
	shirt
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
