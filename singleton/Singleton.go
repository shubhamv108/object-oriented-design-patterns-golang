package singleton

import (
	"fmt"
	"sync"
)

var instance *Singleton
var lock = &sync.Mutex{};

type Singleton struct {
}

func GetInstance() *Singleton {
	if instance == nil {
		lock.Lock();
		defer lock.Unlock();
		if instance == nil {
			fmt.Println("Creating single instance now");
			instance = &Singleton{}
		}
	}
	return instance;
}