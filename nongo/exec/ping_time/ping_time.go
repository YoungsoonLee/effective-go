package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
)

func median(values []float64) float64 {
	// Don't mutate original slice
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)

	i := len(nums) / 2

	if len(nums)%2 == 0 {
		return (nums[i-1] + nums[i]) / 2
	}

	return nums[i]
}

// findTime finds the ping time in ping output line
// returns value, found and error
func findTime(line []byte) (float64, bool, error) {
	// ex, 64 bytes from 161.35.125.80: icmp_seq=0 ttl=55 time=82.634 ms
	var prefix = []byte("time=")
	start := bytes.Index(line, prefix)
	if start == -1 {
		return 0, false, nil
	}

	start += len(prefix) // skip over "time="
	end := bytes.IndexByte(line[start:], ' ')
	if end == -1 {
		return 0, false, fmt.Errorf("can't find end of time value")
	}
	end += start
	val, err := strconv.ParseFloat(string(line[start:end]), 64)
	if err != nil {
		return 0, false, err
	}

	return val, true, nil
}

// medianPing returns the median ping time of the given host
func medianPing(host string, count int) (float64, error) {
	sw := "-c"
	if runtime.GOOS == "windows" {
		sw = "-n" // windows uses -n instead of -c
	}

	// ping -c 4 pragprog.com
	cmd := exec.Command("ping", sw, fmt.Sprintf("%d", count), host)

	data, err := cmd.Output()
	if err != nil { // wait for ping to finish
		return 0, err
	}

	value := make([]float64, 0, count)
	s := bufio.NewScanner(bytes.NewReader(data))
	for s.Scan() {
		//fmt.Println(s.Text())
		val, found, err := findTime(s.Bytes())
		if err != nil {
			return 0, err
		}
		if found {
			value = append(value, val)
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return median(value), nil
}

func main() {
	host := "pragprog.com"
	count := 4

	median, err := medianPing(host, count)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Median ping time for %s over %d pings: %.2f ms\n", host, count, median)
}
