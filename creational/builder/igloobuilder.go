package builder

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.WindowType = "snow"
}
func (b *IglooBuilder) SetDoorType() {
	b.DoorType = "snow"
}
func (b *IglooBuilder) SetNumfloor() {
	b.Floor = 1
}
func (b *IglooBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
