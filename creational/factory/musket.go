package factory

// Concrete product
type Musket struct {
	Gun
}

// GetName implements IGun.
func (m *Musket) GetName() string {
	return m.name
}

// GetPower implements IGun.
func (m *Musket) GetPower() int {
	return m.power
}

// SetName implements IGun.
func (m *Musket) SetName(name string) {
	m.name = name
}

// SetPower implements IGun.
func (m *Musket) SetPower(power int) {
	m.power = power
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "model A",
			power: 5,
		},
	}
}
