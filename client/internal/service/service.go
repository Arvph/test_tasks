package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"tg_client/internal/modules"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// buttons
var menuButtons = tgbotapi.NewReplyKeyboard(
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

func MainService(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to the note bot! Click the button below to create a note.")
				msg.ReplyMarkup = menuButtons //tgbotapi.NewReplyKeyboard(
				//tgbotapi.NewKeyboardButtonRow(
				//	tgbotapi.NewKeyboardButton("Create Note"),
				//),
				//)
				bot.Send(msg)

			case "Create Note":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter your note:")
				bot.Send(msg)
			// case "":

			default:
				// Sending the note to the server
				go sendNoteToServer(update)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Note received and sent to the server!")
				bot.Send(msg)
			}
		}
	}
	return nil
}

func sendNoteToServer(update tgbotapi.Update) {

	note := modules.Note{
		UserID: int(update.Message.Chat.ID),
		Text:   update.Message.Text,
		Status: "In process",
	}
	jsonData, err := json.Marshal(note)
	if err != nil {
		log.Printf("Error marshalling note: %s", err)
		return
	}

	// Change the URL to your server's API endpoint
	resp, err := http.Post("http://localhost:8088/tasks", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending note to server: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected response status: %s", resp.Status)
	}

}

// func Create(msg tgbotapi.MessageConfig) {

// }

// for update := range updates {
// 	if update.Message == nil { // ignore non-Message updates
// 		continue
// 	}

// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

// 	switch update.Message.Text {
// 	case "/start":
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, Bitch!")
// 		bot.Send(msg)
// 	case "open":
// 		msg.ReplyMarkup = menuButtons
// 	case "close":
// 		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
// 	case "Create Note":
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 		// Create(msg)
// 	}

// 	if _, err := bot.Send(msg); err != nil {
// 		log.Panic(err)
// 	}
// }
