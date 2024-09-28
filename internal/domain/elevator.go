package domain

import (
	"sort"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Idle
)

type Elevator struct {
	CurrentFloor int
	Direction    Direction
	Requests     []int
}

func (e *Elevator) AddRequest(floor int) {
	e.Requests = append(e.Requests, floor)
	sort.Ints(e.Requests)
}

func (e *Elevator) MoveElevator() {
	for len(e.Requests) > 0 {
		targetFloor := e.Requests[0]

		if e.CurrentFloor < targetFloor {
			e.Direction = Up
		} else if e.CurrentFloor > targetFloor {
			e.Direction = Down
		}

		for e.CurrentFloor != targetFloor {
			if e.Direction == Up {
				e.CurrentFloor++
			} else {
				e.CurrentFloor--
			}

			time.Sleep(1 * time.Second)
		}

		time.Sleep(4 * time.Second)

		e.Requests = e.Requests[1:]
	}
}
