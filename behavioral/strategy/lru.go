package strategy

import "fmt"

type LRU struct{}

func (l *LRU) Evict(c *Cache) {
	fmt.Println("Evicting by LRU strtegy")
}
