package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	philosophersCount = 5
)

type (
	Eater interface {
		Eat()
	}

	Philosopher struct {
		id          int
		ate         int
		left, right *Chopstick
	}

	Chopstick struct {
		sync.Mutex
		id int
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (p *Philosopher) Eat() {
	p.left.Lock()
	p.right.Lock()

	fmt.Printf("starting to eat %d\n", p.id)
	// nom nom
	fmt.Printf("finishing eating %d\n", p.id)
	p.ate++

	p.right.Unlock()
	p.left.Unlock()
}

func main() {
	fmt.Println("starting dinner")

	// init chopsticks
	var chopsticks []*Chopstick
	for i := 0; i < philosophersCount; i++ {
		c := &Chopstick{
			id: i,
		}
		chopsticks = append(chopsticks, c)
	}

	// init philosophers
	var philosophers []*Philosopher
	for i := 0; i < philosophersCount; i++ {
		p := &Philosopher{
			id:    i + 1,
			left:  chopsticks[i],
			right: chopsticks[(i+1)%philosophersCount],
		}
		philosophers = append(philosophers, p)
	}

	var tickets []int
	id := 1
	tc := philosophersCount * 3
	for i := 1; i <= tc; i++ {
		if i%2 == 1 {
			id -= 1
		} else {
			id += 2
		}
		// fix last ticket
		if i == tc {
			id += 2
		}
		tickets = append(tickets, id%philosophersCount)
	}

	manager := make(chan struct{}, 2)
	var waiter sync.WaitGroup
	for {
		if len(tickets) == 0 {
			break
		}
		// start eating
		manager <- struct{}{}
		// get ticket
		t := tickets[0]
		tickets = tickets[1:]
		waiter.Add(1)
		go func() {
			philosophers[t].Eat()
			<-manager
			waiter.Done()
		}()
	}
	// wait until all stop eating
	waiter.Wait()

	for _, p := range philosophers {
		state := "full"
		if p.ate < 3 {
			state = "hungry"
		}
		if p.ate > 3 {
			state = "overeaten"
		}
		fmt.Printf("i'm %s %d\n", state, p.id)
	}

	fmt.Println("dinner complete, thanks for visiting")
}
