package main

import (
	"bytes"
	"fmt"
)

func genSelect(table string, columns []string) (string, error) {
	var buf bytes.Buffer

	if len(columns) == 0 {
		return "", fmt.Errorf("empty select")
	}
	fmt.Fprintln(&buf, "SELECT ")
	for i, col := range columns {
		suffix := ","
		if i == len(columns)-1 {
			suffix = "" // last column
		}

		fmt.Fprintf(&buf, "    %s%s\n", col, suffix)
	}
	fmt.Fprintf(&buf, "FROM %s;", table)

	return buf.String(), nil
}
