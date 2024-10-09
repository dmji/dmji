package main

import "os"

func main() {
	os.WriteFile("index.html", []byte("hello world"), 0644)
}
