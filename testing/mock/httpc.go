package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient is an API client.
type APIClient struct {
	baseURL string
	c       *http.Client
}

// NewAPIClient returns a new API client.
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		c:       &http.Client{},
	}
}

// Users returns the list of users.
func (c *APIClient) Users() ([]string, error) {
	url := fmt.Sprintf("%s/users", c.baseURL)
	resp, err := c.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []string
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
