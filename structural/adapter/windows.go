package adapter

import "fmt"

type Windows struct{}

func (m *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into mac machine.")
}
