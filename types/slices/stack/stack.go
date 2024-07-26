package main

import (
	"errors"
	"fmt"
)

// Token is a token in a lexer
type Token struct {
	Loc  int
	Char rune
}

// Stack is a stack of tokens
type Stack []Token

// Len returns the number of elements in the stack
func (s Stack) Len() int {
	return len(s)
}

// Push adds a token to the stack
func (s *Stack) Push(t Token) {
	*s = append(*s, t)
}

var (
	// ErrEmpty is returned when the stack is empty
	ErrEmpty = errors.New("empty stack")
)

// Pop removes and returns the top token from the stack
func (s *Stack) Pop() (Token, error) {
	size := s.Len()

	if size == 0 {
		return Token{}, ErrEmpty
	}

	sl := *s // so we won't have to use (*s) all the time
	val := sl[size-1]
	sl = sl[:size-1]

	// if we shrank by more than half and larger than 1k, free memory
	if len(sl) > 1024 && 2*len(sl) < cap(sl) {
		sl2 := make([]Token, len(sl))
		copy(sl2, sl)
		sl = sl2
	}

	*s = sl

	return val, nil
}

func main() {
	var s Stack
	fmt.Println(s)
	s.Push(Token{19, '('})
	s.Push(Token{49, '['})
	fmt.Println(s)

	if v, err := s.Pop(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("pop", v)
	}
}
