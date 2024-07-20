package main

import (
	"fmt"
	"reflect"
	"time"
)

// Log is a log event.
type Log struct {
	Time  time.Time `db:"ts"`
	Level int       `db:"level"`
	Text  string    `db:"message"`
}

func parseStructTags(s any) (map[string]string, error) {
	typ := reflect.TypeOf(s)
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is not a struct", typ)
	}

	m := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		if dbName := fld.Tag.Get("db"); dbName != "" {
			m[fld.Name] = dbName
		}
	}

	return m, nil
}

func main() {
	var l Log
	fmt.Println(parseStructTags(l))

	var n int
	fmt.Println(parseStructTags(n))
}
