package decorator

type CheeseTopping struct {
	Pizza IPizza
}

func (t *CheeseTopping) GetPrice() int {
	return t.GetPrice() + 5
}
