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
	if ok, _ := exists("_site"); !ok {
		os.Mkdir("_site", 0700)
	}

	os.WriteFile("_site/index.html", []byte("Hello from actions!"), 0644)
}
