package main

import "encoding/json"

// Stack is a LIFO stack.
type Stack struct {
	Value string
	Next  *Stack
}

// MarshalJSON implements the json.Marshaler interface.
func (s *Stack) MarshalJSON() ([]byte, error) {
	var values []string
	for s != nil {
		values = append(values, s.Value)
		s = s.Next
	}

	return json.Marshal(values)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Stack) UnmarshalJSON(data []byte) error {
	var values []string
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	var head *Stack
	for i := len(values) - 1; i >= 0; i-- {
		head = &Stack{Value: values[i], Next: head}
	}

	*s = *head
	return nil
}
