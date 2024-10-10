package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	params := &bot.SendMessageParams{
		ChatID: os.Getenv("TELEGRAM_ID_CHAN"),
		Text:   time.Now().String(),
	}

	b.SendMessage(ctx, params)

}
