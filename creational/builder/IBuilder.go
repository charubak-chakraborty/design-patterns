package builder

// Builder is a creational design pattern, which allows constructing complex objects step by step.
// Unlike other creational patterns, Builder doesn’t require products to have a common interface. That makes it possible to produce different products using the same construction process.

// The Builder pattern is used when the desired product is complex and requires multiple steps to complete. In this case, several construction methods would be simpler than a single monstrous constructor. The potential problem with the multistage building process is that a partially built and unstable product may be exposed to the client. The Builder pattern keeps the product private until it’s fully built.

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumfloor()
	GetHouse() House
}
type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return NewNormalBuilder()
	} else if builderType == "igloo" {
		return NewIglooBuilder()
	}
	return nil
}
