package bot

import (
	"classconnect-bot/internal/config"
	"classconnect-bot/internal/handler"
	"classconnect-bot/internal/service"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type Bot struct {
	cfg     *config.Config
	bot     *telebot.Bot
	handler *handler.Handler
}

func New() *Bot {
	cfg := config.MustLoad()

	pref := telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	classConnectService := service.NewService(cfg)

	handlers := handler.NewHandler(bot, classConnectService)

	return &Bot{
		cfg:     cfg,
		bot:     bot,
		handler: handlers,
	}
}

func (b *Bot) Run() {
	b.bot.Handle("/start", b.handler.StartHandler)
	b.bot.Handle("/groups", b.handler.ShowGroupsHandler)
	b.bot.Handle("/join", b.handler.JoinToGroupHandler)

	b.bot.Handle("👥 My group", b.handler.GetGroupInfoHandler)
	b.bot.Handle("🗓️ Getting a schedule for today", b.handler.GetScheduleForDayHandler)
	b.bot.Handle("❌ Exit", b.handler.ExitFromGroupHandler)

	//b.bot.Handle("Show groups", b.handler.ShowGroupsHandler)
	b.bot.Handle(telebot.OnCallback, b.handler.CallbackHandler)

	b.bot.Start()
}
