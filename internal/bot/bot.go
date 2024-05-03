package bot

import "classconnect-bot/internal/config"

type Bot struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Bot {
	return &Bot{
		cfg: cfg,
	}
}

func (b *Bot) Run() {

}
