package main

import (
	"fmt"
	"sync"

	"github.com/arthurbaquit/golang-studies/rabbitmq/consumer"
	"github.com/arthurbaquit/golang-studies/rabbitmq/publisher"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	publisher.DeclareQueue("testing")
	go consumer.ConsumeFromQueue("testing")
	for i := 0; i < 10; i++ {
		publisher.PublishIntoQueue(fmt.Sprintf("testing message number: %v", i), "testing")
	}
	wg.Add(1)
	go func() {
		for i := 0; i < 100000; i++ {
			// Await to garantee that the consumer will receive all messages
		}
		wg.Done()
	}()
	wg.Wait()

}
