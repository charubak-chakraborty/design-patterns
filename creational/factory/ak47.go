package factory

// Concrete product
type AK47 struct {
	Gun
}

// GetName implements IGun.
func (a *AK47) GetName() string {
	return a.name
}

// GetPower implements IGun.
func (a *AK47) GetPower() int {
	return a.power
}

// SetName implements IGun.
func (a *AK47) SetName(name string) {
	a.name = name
}

// SetPower implements IGun.
func (a *AK47) SetPower(power int) {
	a.power = power
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "model 1",
			power: 4,
		},
	}
}
