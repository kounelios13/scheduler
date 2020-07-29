package scheduler

import (
	"errors"
	"fmt"
	"sync"
)

// Scheduler struct
type Scheduler struct {
	Name  string
	Tasks []interface{}
	m     sync.Mutex
}

// NewScheduler Create an empty scheduler
func NewScheduler(name string) (*Scheduler, error) {
	if name == "" {
		return nil, errors.New("scheduler name is empty")
	}
	return &Scheduler{
		Name:  name,
		Tasks: make([]interface{}, 0),
	}, nil
}

// SayHi Say hi
func (s *Scheduler) SayHi() {
	fmt.Printf("Hello .This is scheduler %s talking \n", s.Name)
}

// AddTask Add a task
func (s *Scheduler) AddTask(task interface{}) {
	s.m.Lock()
	s.Tasks = append(s.Tasks, task)
	s.m.Unlock()
}
