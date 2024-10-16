package abstractFactory

type Nike struct{}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "",
			size: 0,
		},
	}
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "",
			size: 0,
		},
	}
}
