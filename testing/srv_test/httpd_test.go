package srvtest

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"
)

func buildServer(t *testing.T) string {
	fileName := path.Join(t.TempDir(), "httpd")
	cmd := exec.Command("go", "build", "-o", fileName, "httpd.go")
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build server: %v", err)
	}
	return fileName
}

func freePort(t *testing.T) int {
	conn, err := net.Listen("tcp", "")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	conn.Close()
	return conn.Addr().(*net.TCPAddr).Port
}

func waitForServer(t *testing.T, addr string) {
	start := time.Now()
	timeout := 10 * time.Second
	var err error
	var conn net.Conn
	for time.Since(start) < timeout {
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			conn.Close()
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("server did not start in %v: %v", timeout, err)
}

func runServer(t *testing.T) int {
	exe := buildServer(t)
	t.Logf("Server exec: %q", exe)
	port := freePort(t)
	t.Logf("Server port: %d", port)

	env := os.Environ()
	env = append(env, fmt.Sprintf("HTTPD_ADDR=:%d", port))
	cmd := exec.Command(exe)
	cmd.Env = env
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("localhost:%d", port)
	waitForServer(t, addr)
	t.Cleanup(func() {
		if err := cmd.Process.Kill(); err != nil {
			t.Logf("warning: can't kill server (pid=%d)", cmd.Process.Pid)
		}
	})

	return port
}

func TestHealth(t *testing.T) {
	port := runServer(t)

	url := fmt.Sprintf("http://localhost:%d/health", port)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %q: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET %q: status code = %d, want %d", url, resp.StatusCode, http.StatusOK)
	}
}
