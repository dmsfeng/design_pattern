package singleton

import (
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var singleInstance *singleton

func getInstance() *singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {

			singleInstance = &singleton{}
		}
	}

	return singleInstance
}
