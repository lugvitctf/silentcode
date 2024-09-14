package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"
)

// In the code where darkness drifts,
// Two paths are masked with clever twists.
// One will guide you to the start,
// The other reveals a cryptic part.
// ⁣ (https://t.me/stuxn3tbot?start=1mas3cr3t) ⁣ (https://t.me/stuxn3tbot?start=t0n3xts13p)
// Search the depths where silence reigns,
// And uncover where the key remains.
// The hidden links hold the way,
// To lead you through this darkened play.

// To find the key that hides from view,
// Shift letters round to uncover the clue.
// In a dance of characters, take a step,
// And the hidden message will be adept. ⁣ (https://t.me/stuxn3tbot?start=f0ld3ds3c)

// 	940fabb3cc45db6a

// Through iron walls of cipher's might,
// One hundred twenty-eight threads hold tight.
// A dance of locks, a code unbound,
// Stuxnet whispers without a sound.
// ⁣ (https://t.me/stuxn3tbot?start=n0td0n3y3t)
// In echoes of a hidden scheme,
// He passes the flag, unseen, supreme.
// A fortress built on silent creed,
// An ancient cipher, swift with speed.

func main() {
	_ = godotenv.Load()
	err := StartDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey != "" {
		API_KEY = apiKey
		log.Println("Loaded API Key:", apiKey)
	}
	bot, err := gotgbot.NewBot(os.Getenv("BOT_TOKEN"), nil)
	if err != nil {
		panic(err)
	}
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", start))

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: false,
	})
	if err != nil {
		panic(err)
	}
	log.Println("Bot has been started")
	mux := http.NewServeMux()
	mux.HandleFunc("/new", newFlag(bot.Username))
	fmt.Println(http.ListenAndServe(":4853", mux))
}
