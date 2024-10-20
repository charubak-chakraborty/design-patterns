package observer

import "fmt"

type Customer struct{ ID string }

func (c *Customer) Update(item string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, item)

}
func (c *Customer) GetID() string {
	return c.ID
}
