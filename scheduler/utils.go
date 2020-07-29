package scheduler

import (
	"errors"
	"fmt"
)

// Scheduler struct
type Scheduler struct {
	Name string
}

// NewScheduler Create an empty scheduler
func NewScheduler(name string) (*Scheduler, error) {
	if name == "" {
		return nil, errors.New("scheduler name is empty")
	}
	return &Scheduler{
		Name: name,
	}, nil
}

// SayHi Say hi
func (s Scheduler) SayHi() {
	fmt.Printf("Hello .This is scheduler %s talking \n", s.Name)
}
