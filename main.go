package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/dmji/dmji/components"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func init() {

	path := ".env"
	for i := range 10 {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}
	panic(".env not found")
}

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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramIdChan := os.Getenv("TELEGRAM_ID_CHAN")

	if ok, _ := exists("_site"); !ok {
		os.Mkdir("_site", 0700)
	}

	/*
	   resp, err := http.Get("https://animelayer.ru/")

	   	if err != nil {
	   		return
	   	}
	   	defer resp.Body.Close()

	   	bodyBytes, err := io.ReadAll(resp.Body)
	   	if err != nil {
	   		return
	   	}
	*/

	f, err := os.Create("_site/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	components.IndexPage(time.Now()).Render(ctx, f)

	/*
		p := animelayer.New(&http.Client{})
		cats, err := p.CategoryPageToPartialItems(ctx, animelayer.Categories.Anime(), 1)
		if err != nil {
			panic(err)
		}

		t := item.ItemInfoSet{}
		items := make([]animelayer.ItemPartial, 0, len(cats))
		for _, c := range cats {
			if c.Error == nil {
				items = append(items, *c.Item)
				t.Data = append(t.Data,
					&item.ItemInfo{
						Identifier:  c.Item.Identifier,
						IsCompleted: c.Item.IsCompleted,
						Title:       c.Item.Title,
					})
			}
		}

		byteItems, err := json.Marshal(&items)
		if err != nil {
			panic(err)
		}
		writeZip("test_gzip_json", byteItems)

		protoItems, err := proto.Marshal(&t)
		if err != nil {
			panic(err)
		}
		writeZip("test_gzip_proto", protoItems)
		return
	*/

	opts := []bot.Option{}

	b, err := bot.New(telegramToken, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	params := &bot.SendMessageParams{
		ChatID: telegramIdChan,
		Text:   "Ping!",
	}

	b.SendMessage(ctx, params)
	/*
		info, err := b.GetChat(ctx, &bot.GetChatParams{ChatID: telegramIdChan})
		if nil != err {
			panic(err)
		}

		_ = info
	*/

	/*
		 	b, err := telego.NewBot(telegramToken)
			if err != nil {
				panic(err)
			}

			b.GetChat(&telego.GetChatParams{})
	*/
}
