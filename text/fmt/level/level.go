package main

import "fmt"

// Priority is bug priority.
type Priority uint8

const (
	// Low is low priority.
	Low Priority = 10
	// Medium is medium priority.
	Medium Priority = 20
	// High is high priority.
	High Priority = 30
)

// String implements fmt.Stringer interface.
func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	}

	return fmt.Sprintf("<%d>", p)
}

// Bug is a struct that represents a bug.
type Bug struct {
	Title    string
	Priority Priority
}

func main() {
	bug := &Bug{"Crash on startup", High}
	fmt.Printf("bug: %+v\n", bug)

}
