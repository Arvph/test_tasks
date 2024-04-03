package telegram

import (
	"log"
	"tg_client/internal/config"
	"tg_client/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Create Note"),
		tgbotapi.NewKeyboardButton("Update Note"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get all notes"),
		tgbotapi.NewKeyboardButton("Get Note"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Delete Note"),
	),
)

func StartBot(token config.Config) error {
	// здесь формируем код телеграм бота
	log.Println("TG is starting")

	bot, err := tgbotapi.NewBotAPI(token.Token)
	if err != nil {
		log.Fatal("unable to crate TG bot", err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	if err = service.MainService(bot, updates); err != nil {
		log.Println(err)
	}

	log.Println("tg is running")

	return nil
}

// func request(update tgbotapi.Update) error {

// }

// надо создать запрос на подключение. Если ID поьзователя не известно, то создать запись в БД (при нажатии на /start должен создаваться новый клиент)
// если такой клиент уже есть предложить кнопки по создания, чтонию, удалению и изменеию заметок.

// if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
// 	switch update.Message.Text {
// 	case "/start":
// 		log.Println("user_id", update.Message.From.ID)
// 		// log.Println("user_id", update.Message.Chat.ID)
// 		log.Println("user_id", update.Message.From.UserName)
// 		// log.Println("user_id", update.Message.Chat.ID)

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, Bitch!")
// 		bot.Send(msg)

// 	default:
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 		bot.Send(msg)
// 	}

// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// bot.Send(msg)
