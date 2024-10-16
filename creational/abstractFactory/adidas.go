package abstractFactory

type Adidas struct {
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "",
			size: 0,
		},
	}
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "",
			size: 0,
		},
	}
}
