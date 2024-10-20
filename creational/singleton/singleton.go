package singleton

import "sync"

var lock = &sync.Mutex{}

type Singleton struct{}

var singleton *Singleton

// Singleton is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.

// There is a nil-check at the start for making sure singleInstance is empty first time around. This is to prevent expensive lock operations every time the getinstance method is called. If this check fails, then it means that the singleInstance field is already populated.

// The singleInstance struct is created within the lock.

// There is another nil-check after the lock is acquired. This is to ensure that if more than one goroutine bypasses the first check, only one goroutine can create the singleton instance. Otherwise, all goroutines will create their own instances of the singleton struct.

func GetInstance() *Singleton {

	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			singleton = &Singleton{}
		}
	}

	return singleton
}
