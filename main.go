package main

import (
	"io"
	"net/http"
	"os"
)

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

	resp, err := http.Get("https://animelayer.ru/")

	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	os.WriteFile("_site/index.html", bodyBytes, 0644)
}
