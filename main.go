package main

import (
	"fmt"

	"github.com/charubak.chakraborty/design-patterns-golang/creational/abstractFactory"
	"github.com/charubak.chakraborty/design-patterns-golang/creational/factory"
)

func main() {

	//creational - factory
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")
	ak47.GetName()
	musket.GetName()

	//creational - abstract factory
	adidasFactory, _ := abstractFactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportsFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()
	adidasShoe.SetLogo("logo 1")
	fmt.Print(adidasShoe.GetLogo())
	adidasShirt.SetLogo("logo 1")
	fmt.Print(adidasShirt.GetLogo())

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()
	nikeShoe.SetLogo("logo 1")
	fmt.Print(nikeShoe.GetLogo())
	nikeShirt.SetLogo("logo 1")
	fmt.Print(nikeShirt.GetLogo())

}
