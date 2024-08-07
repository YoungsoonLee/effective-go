package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockTransport struct {
	body []byte
	err  error
}

// RoundTrip implements http.RoundTripper interface.
func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.err != nil {
		return nil, t.err
	}

	w := httptest.NewRecorder()
	if t.body != nil {
		w.Write(t.body)
	}
	return w.Result(), nil
}

func TestUsersOK(t *testing.T) {
	users := []string{"clark", "bruce", "diana"}
	data, err := json.Marshal(users)
	require.NoError(t, err, "encode JSON")

	c := NewAPIClient("http://localhost:8080")
	c.c.Transport = &MockTransport{data, nil}
	reply, err := c.Users()
	require.NoError(t, err, "fetch users")
	require.Equal(t, users, reply, "users")
}

func TestUsersConnecntionError(t *testing.T) {
	c := NewAPIClient("http://localhost:8080")
	c.c.Transport = &MockTransport{nil, fmt.Errorf("network error")}
	_, err := c.Users()
	require.Error(t, err)
}

func TestUsersBadJSON(t *testing.T) {
	c := NewAPIClient("http://localhost:8080")
	c.c.Transport = &MockTransport{[]byte(`["clark","diana","bruce"`), nil}
	_, err := c.Users()
	require.Error(t, err)
}
