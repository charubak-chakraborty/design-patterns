package main

import (
	"fmt"

	"github.com/charubak.chakraborty/design-patterns-golang/behavioral/observer"
	"github.com/charubak.chakraborty/design-patterns-golang/behavioral/proxy"
	"github.com/charubak.chakraborty/design-patterns-golang/behavioral/strategy"
	"github.com/charubak.chakraborty/design-patterns-golang/creational/abstractFactory"
	"github.com/charubak.chakraborty/design-patterns-golang/creational/builder"
	"github.com/charubak.chakraborty/design-patterns-golang/creational/factory"
	"github.com/charubak.chakraborty/design-patterns-golang/creational/singleton"
	"github.com/charubak.chakraborty/design-patterns-golang/structural/adapter"
	"github.com/charubak.chakraborty/design-patterns-golang/structural/composite"
	"github.com/charubak.chakraborty/design-patterns-golang/structural/decorator"
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

	//creational - singleton
	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()

	//creational - builder
	normalBuilder := builder.GetBuilder("normal")
	igloolBuilder := builder.GetBuilder("iglool")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(igloolBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("igloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("igloo House Num Floor: %d\n", iglooHouse.Floor)

	//structural - adapter

	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &adapter.Windows{}
	windowsAdapter := &adapter.WindowsAdapter{WindowsMachine: windows}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)

	//structural - composite

	file1 := &composite.File{Name: "File1"}
	file2 := &composite.File{Name: "File2"}
	file3 := &composite.File{Name: "File3"}

	folder1 := &composite.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &composite.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")

	//structural - decorator
	pizza := &decorator.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())

	//behavioral - obsevrer
	shirtItem := observer.NewItem("Nike Shirt")

	observerFirst := &observer.Customer{ID: "abc@gmail.com"}
	observerSecond := &observer.Customer{ID: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()

	//behavioural - strategy
	lfu := &strategy.LFU{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &strategy.LRU{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &strategy.FIFO{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

	//structural - proxy
	nginxServer := proxy.NewNGINXServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
