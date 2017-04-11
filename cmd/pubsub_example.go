package main

import (
	"fmt"

	"github.com/jackric/pubsub"
)

type ToyStatefulSubscriber struct {
	seen int
}

func (s *ToyStatefulSubscriber) callback(message interface{}) bool {
	s.seen++
	fmt.Printf("ToyStatefulSubscriber got: %v\n", message)
	if s.seen >= 2 {
		fmt.Println("ToyStatefulSubscriber is now unsubscribing\n")
		return false
	}
	return true
}

func main() {
	b := pubsub.NewBroker(5)
	var err error

	b.Publish("You shouldn't see me")
	err = b.Subscribe(func(message interface{}) bool {
		fmt.Printf("Subscriber got: %v\n", message)
		return true
	})
	if err != nil {
		panic(err)
	}
	toy := ToyStatefulSubscriber{}
	err = b.Subscribe(toy.callback)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		b.Publish(fmt.Sprintf("numero %d", i))
	}

}
