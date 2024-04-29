package singleton

import "sync"

type SyncSingleton struct {
}

var (
	instances *SyncSingleton
	once      sync.Once
)

func GetInstanceSync() *SyncSingleton {
	once.Do(func() {
		instances = &SyncSingleton{}
	})
	return instances
}
