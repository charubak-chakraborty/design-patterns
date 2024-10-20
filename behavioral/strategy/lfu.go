package strategy

import "fmt"

type LFU struct{}

func (l *LFU) Evict(c *Cache) {
	fmt.Println("Evicting by LFU strtegy")
}
