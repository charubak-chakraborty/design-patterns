package strategy

import "fmt"

type FIFO struct{}

func (l *FIFO) Evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}
