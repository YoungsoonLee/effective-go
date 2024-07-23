package main

import "log"

// State is task state
type State byte

const (
	// Ready is the default state
	Ready State = iota + 1
	// Wokring is the state when the task is being worked on
	Wokring
	// Done is the state when the task is done
	Done
)

// Task is a task to execute
type Task struct {
	ID     uint
	Result any
	Err    error
	State  State

	Work    func() (any, error)
	Watches []func(*Task)
}

// NewTask returns a new Task will excute work.
func NewTask(id uint, work func() (any, error)) *Task {
	return &Task{
		ID:    id,
		State: Ready,
		Work:  work,
	}
}

// Subcribe adds a watch function to the task.
func (t *Task) Subcribe(w func(*Task)) {
	t.Watches = append(t.Watches, w)
}

// Execute executes the task.
func (t *Task) Execute() error {
	t.State = Wokring
	defer func() { t.State = Done }()

	t.Result, t.Err = t.Work()
	if t.Err != nil {
		log.Printf("task %d failed: %s", t.ID, t.Err)
	}

	// Notify watches
	for _, w := range t.Watches {
		w(t)
	}

	return t.Err
}

// Watcher is a task watcher
type Watcher struct{}

// Handle handles a task.
func (w *Watcher) Handle(t *Task) {
	log.Printf("task %d is %v", t.ID, t.State)
}

func main() {
	t := NewTask(7, func() (any, error) { return "done", nil })
	t.Subcribe(func(t *Task) {
		log.Printf("info: w1: from %d - %#v %v", t.ID, t.Result, t.Err)
	})

	var w Watcher
	t.Subcribe(w.Handle)

	t.Execute()
}
