package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct{}

var singleInstance *Single

func GetSingleInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Print("single instance creating")
			singleInstance = &Single{}
		} else {
			fmt.Print("single instance already created")
		}
	}
	fmt.Print("single instance already created")
	return singleInstance
}
