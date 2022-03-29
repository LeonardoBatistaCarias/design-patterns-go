package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton map[string]string

var (
	instance singleton
)

func GetInstance() singleton {

	fmt.Println("Creating single instance now.")

	once.Do(func() {
		instance = make(singleton)
		fmt.Println("Created")
	})

	return instance
}
