package builder

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.WindowType = "wood"
}
func (b *NormalBuilder) SetDoorType() {
	b.DoorType = "wood"
}
func (b *NormalBuilder) SetNumfloor() {
	b.Floor = 2
}
func (b *NormalBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
