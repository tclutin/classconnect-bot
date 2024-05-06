package handler

import (
	"classconnect-bot/internal/service"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	bot     *telebot.Bot
	service *service.Service
}

func NewHandler(bot *telebot.Bot, service *service.Service) *Handler {
	return &Handler{
		bot:     bot,
		service: service,
	}
}
