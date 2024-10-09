package main

import "os"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	if ok, _ := exists("www"); !ok {
		os.Mkdir("www", 0700)
	}

	os.WriteFile("www/index.html", []byte("Hello from actions!"), 0644)
	os.WriteFile("www/hello.html", []byte("Hello from actions 2!"), 0644)
}
