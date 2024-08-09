//go:build ignore

package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func run(args ...string) (string, error) {
    cmd := exec.Command(args[0], args[1:]...)
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)), nil
}


func main() {
    version, err := run("git", "tag", "--points-at", "HEAD")
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    if version == "" {
        var err error
        version, err = run("git", "rev-parse", "--short", "HEAD")
        if err != nil {
            log.Fatalf("error: %s", err)
        }
    }

    log.Printf("building version %s\n", version)
    ldflags := fmt.Sprintf("-ldflags=-X main.version=%s", version)
    _, err = run("go", "build", ldflags, "-o", "agent")
    if err != nil {
        log.Fatalf("error: %s", err)
    }
}
