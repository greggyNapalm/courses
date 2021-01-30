package main

import (
	"fmt"
	"sync"
	"time"
)

type Host struct {
	eatingPhilsState [5]int
	eatingPhilsNum   int
	eatenPhilsNum    int
	requestsBacklog  []int
	requsts, reports chan int
	permits          [5]chan bool
}
type ChopS struct{ sync.Mutex }
type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat(numOfEats int, cToRequest, cToReport chan int, cToGetPermit chan bool) {
	for i := 0; i < numOfEats; i++ {
		cToRequest <- p.number
		<-cToGetPermit
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat ", p.number+1)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("finishing eating ", p.number+1)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		cToReport <- p.number
	}
}

func (h *Host) Run(numOfPhils, numOfEatsForPhil int) {
	for {
		select {
		case requesterId := <-h.requsts:
			if h.isEeligibleToEat(requesterId) {
				h.allowToEat(requesterId)
			} else {
				h.requestsBacklog = append(h.requestsBacklog, requesterId)
			}
		case reporterId := <-h.reports:
			h.eatingPhilsState[reporterId] = 0
			h.eatingPhilsNum--
			h.eatenPhilsNum++
		}
		h.walkBacklog()
		if h.eatenPhilsNum == numOfPhils*numOfEatsForPhil {
			fmt.Println("*All philosophers ate successfully*")
			break
		}
	}
}
func (h *Host) isEeligibleToEat(reqId int) bool {
	if h.eatingPhilsNum == 2 {
		return false
	}
	if h.eatingPhilsState[(reqId+4)%5] != 0 || h.eatingPhilsState[(reqId+1)%5] != 0 {
		// Not one of the neighbors should eat.
		// This means that both required sticks are free.
		return false
	}
	return true
}
func (h *Host) allowToEat(reqId int) {
	h.eatingPhilsState[reqId] = 1
	h.eatingPhilsNum++
	h.permits[reqId] <- true
}
func (h *Host) walkBacklog() {
	eligableWasFound := false
	for i := 0; i < len(h.requestsBacklog); i++ {
		if h.isEeligibleToEat(h.requestsBacklog[i]) {
			h.allowToEat(h.requestsBacklog[i])
			h.requestsBacklog = append(h.requestsBacklog[:i], h.requestsBacklog[i+1:]...)
			eligableWasFound = true
			break
		}
	}
	if eligableWasFound == true {
		h.walkBacklog()
	}
}
func main() {
	numOfPhils := 5
	numOfEatsForPhil := 3
	cRequests := make(chan int)
	cReports := make(chan int)
	var cPermits [5]chan bool
	CSticks := make([]*ChopS, 5)
	philos := make([]Philo, numOfPhils)
	for i := 0; i < numOfPhils; i++ {
		CSticks[i] = new(ChopS)
	}
	for i := 0; i < numOfPhils; i++ {
		cPermits[i] = make(chan bool)
		philos[i] = Philo{CSticks[i], CSticks[(i+1)%5], i}
	}
	for i := 0; i < numOfPhils; i++ {
		go philos[i].eat(numOfEatsForPhil, cRequests, cReports, cPermits[i])
	}
	h := Host{
		eatingPhilsState: [5]int{},
		eatingPhilsNum:   0,
		requestsBacklog:  []int{},
		requsts:          cRequests,
		reports:          cReports,
		permits:          cPermits,
	}
	h.Run(numOfPhils, numOfEatsForPhil)
}
