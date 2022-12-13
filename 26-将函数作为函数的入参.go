package main

import "fmt"

type Callback func() string

func Load(key string, callback Callback) string {
	if key == "" {
		return callback()
	} else {
		return key
	}
}

func main() {
	value := Load("", func() string {
		return "empty string"
	})
	fmt.Println("cache: ", value)
}
