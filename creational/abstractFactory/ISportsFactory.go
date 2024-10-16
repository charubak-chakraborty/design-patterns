package abstractFactory

import "errors"

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	} else if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, errors.New("wrong brand")
}
