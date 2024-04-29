package singleton

import (
	"sync"
)

type DoubleSingleton struct{}

var (
	instanceDouble *DoubleSingleton
	onces          sync.Once
	mu             sync.Mutex
)

func GetInstanceDouble() *DoubleSingleton {
	if instanceDouble == nil {
		mu.Lock()
		defer mu.Unlock()
		instanceDouble = &DoubleSingleton{}
	}
	return instanceDouble
}
