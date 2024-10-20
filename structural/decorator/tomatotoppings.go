package decorator

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.GetPrice() + 2
}
