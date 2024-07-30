package main

import "encoding/json"

// UserRequest is a user request.
type UserRequest struct {
	Login string
}

// GroupRequest is a request for a group.
type GroupRequest struct {
	ID string
}

// Request is the set of all possible requests.
type Request interface {
	UserRequest | GroupRequest
}

// UnmarshalJSON implements the json.Unmarshaler.
func UnmarshalJSON[T Request](data []byte, d *T) error {
	return json.Unmarshal(data, d)
}

func main() {

}
