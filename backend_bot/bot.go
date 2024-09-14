package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/fizzywhizbang/gorot"
)

var (
	l0poem1 = `In the code where darkness drifts,
Two paths are masked with clever twists.
One will guide you to the start,
The other reveals a cryptic part.`
	l0poem2 = `Search the depths where silence reigns,
And uncover where the key remains.
The hidden links hold the way,
To lead you through this darkened play.`
)

var (
	l1_1poem = `To find the key that hides from view,
Shift letters round to uncover the clue.
In a dance of characters, take a step,
And the hidden message will be adept.`

	l1_2poem1 = `Through iron walls of cipher's might,
One hundred twenty-eight threads hold tight.
A dance of locks, a code unbound,
Stuxnet whispers without a sound.`

	l1_2poem2 = `In echoes of a hidden scheme,
He passes the flag, unseen, supreme.
A fortress built on silent creed,
An ancient cipher, swift with speed.`
)

func respStart(bot *gotgbot.Bot, ctx *ext.Context) error {
	text := fmt.Sprintf(`%s
%s %s
%s`, l0poem1,
		"<a href=\"https://t.me/"+bot.Username+"?start=1mas3cr3t\">\u2063</a>",
		"<a href=\"https://t.me/"+bot.Username+"?start=t0n3xts13p\">\u2063</a>",
		l0poem2,
	)
	ctx.EffectiveMessage.Reply(bot, text, &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeHTML,
	})
	return ext.EndGroups
}

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()
	user := ctx.EffectiveUser
	if user == nil {
		return ext.EndGroups
	}
	if len(args) > 1 {
		if strings.HasPrefix(args[1], "n_") {
			sessionId := strings.TrimPrefix(args[1], "n_")
			session := GetSession(sessionId)
			if session == nil {
				return ext.EndGroups
			}
			AddUserToSession(sessionId, user.Id)
			return respStart(bot, ctx)
		}
		sessionId := GetUserSession(user.Id)
		if sessionId == "" {
			return ext.EndGroups
		}
		session := GetSession(sessionId)
		switch args[1] {
		case "1mas3cr3t":
			_, err := ctx.EffectiveMessage.Reply(bot, fmt.Sprintf(`%s %s`,
				l1_1poem,
				"<a href=\"https://t.me/"+bot.Username+"?start=f0ld3ds3c\">\u2063</a>",
			), &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeHTML,
			})
			if err != nil {
				log.Println(err.Error())
			}
		case "f0ld3ds3c":
			_, err := ctx.EffectiveMessage.Reply(bot, gorot.Encode("Hello from Stuxnet! Here is the secret: "+session.Password), nil)
			if err != nil {
				log.Println(err.Error())
			}
		case "t0n3xts13p":
			text := fmt.Sprintf("%s\n<a href=\"https://t.me/%s?start=n0td0n3y3t\">\u2063</a>\n%s", l1_2poem1, bot.Username, l1_2poem2)
			_, err := ctx.EffectiveMessage.Reply(bot, text, &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeHTML,
			})
			if err != nil {
				log.Println(err.Error())
			}
		case "n0td0n3y3t":
			_, err := ctx.EffectiveMessage.Reply(bot, session.Flag, nil)
			if err != nil {
				log.Println(err.Error())
			}
		default:
			return ext.EndGroups
		}
		return ext.EndGroups
	}
	sessionId := GetUserSession(user.Id)
	if sessionId == "" {
		return ext.EndGroups
	}
	return respStart(bot, ctx)
}
