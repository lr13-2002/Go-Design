package singleton

import "sync"

var once sync.Once

func GetLazyInstance() *Singleton {
	if singleton == nil {
		once.Do(func() {
			singleton = new(Singleton)
		})
	}
	return singleton
}
