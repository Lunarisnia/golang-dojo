package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jake := Student{"Jake"}
	julia := Student{"Julia"}
	for i := 0; i < 10; i++ {
		raisedHandChannel := <-fanIn(jake.RaiseHand(), julia.RaiseHand())
		if raisedHandChannel == nil {
			fmt.Println("Nobody answer..., Question expired.")
		} else {
			fmt.Printf("%v, raised their hand and provided an answer. It was %v.\n", raisedHandChannel, raisedHandChannel.ProvideAnswer())
		}
	}
}

func fanIn(p1, p2 <-chan Agent) <-chan Agent {
	queueChannel := make(chan Agent)
	go func() {
		for {
			select {
			case a := <-p1:
				queueChannel <- a
			case a := <-p2:
				queueChannel <- a
			case <-time.After(1 * time.Second):
				queueChannel <- nil
			}
		}
	}()
	return queueChannel
}

// func raiseHand(name string) <-chan string {
// 	ch := make(chan string)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 			ch <- name
// 		}
// 	}()
// 	return ch
// }

type Agent interface {
	RaiseHand() <-chan Agent
	ProvideAnswer() string
	String() string
}

type Student struct {
	name string
}

func (s Student) ProvideAnswer() string {
	if rand.Intn(100) < 50 {
		return "CORRECT"
	}
	return "WRONG"
}

func (s Student) RaiseHand() <-chan Agent {
	ch := make(chan Agent)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(5e3)) * time.Millisecond)
			ch <- s
		}
	}()
	return ch
}

func (s Student) String() string {
	return s.name
}
