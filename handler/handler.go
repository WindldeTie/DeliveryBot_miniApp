package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strings"
)

type Handler struct {
	bot *tgbotapi.BotAPI
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		bot: bot,
	}
}

func (h *Handler) Start(debug bool) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	h.bot.Debug = debug
	updates := h.bot.GetUpdatesChan(u)
	go h.console()

	for update := range updates {
		h.HandleUpdate(update)
	}
}

// Обработка команд --------------------------------------------------------------------------------------------------

func (h *Handler) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		command := strings.TrimSpace(update.Message.Text)
		switch command {
		case "/start":
			h.handleStart(update)
			return
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			h.bot.Send(msg)
		}
	}
	if update.CallbackQuery != nil {

	}
}

// Вспомогательные функции --------------------------------------------------------------------------------------------

func (h *Handler) handleStart(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Чтобы сделать заказ, нажмите на кнопку снизу")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сделать Заказ", "/order"),
		),
	)
	h.bot.Send(msg)
}

func (h *Handler) handelOrder(callback *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Нажмите на кнопку снизу и сделайте заказ!")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Сделать заказ", os.Getenv("webAppURL")),
		),
	)
	h.bot.Send(msg)
}

func (h *Handler) callbackOrder(callback *tgbotapi.CallbackQuery) {
	data := callback.Data
	if data == "/order" {
		h.handelOrder(callback)
	}
}
