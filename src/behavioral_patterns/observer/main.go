package main

import "fmt"

func main() {
	observable := NewsAgency{}
	observer := NewsChannel{}

	observable.AddObserver(&observer)
	observable.SetNews("news")
	fmt.Println(observer.GetNews())
}
