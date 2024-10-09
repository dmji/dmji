package main

import "os"

func main() {
	os.WriteFile("index.html", []byte("Hello from actions!"), 0644)
}
