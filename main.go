package main

import (
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botToken = "6796961656:AAGimXMVJzd0a1JwkFvSEqR28mbMQr2aL1k"
)

// Структура для отслеживания состояний
type UserState struct {
	CurrentState string
	PrevState    string
}

var (
	bot             *tgbotapi.BotAPI
	userStates      = make(map[int64]*UserState)
	userStatesMutex sync.Mutex
	mainMenu        = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Описать проблему"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Наши услуги!"),
			tgbotapi.NewKeyboardButton("О нас!"),
			tgbotapi.NewKeyboardButton("Предложения!"),
		),
	)
	describeProblemMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Не работает"),
			tgbotapi.NewKeyboardButton("Проблемы с оплатой"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Другая проблема"),
			tgbotapi.NewKeyboardButton("Вернуться в предыдущее меню"),
		),
	)
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handleUpdate(update)
	}
}

func handleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		handleMessage(update.Message)
	}
}

func handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	switch message.Text {
	case "/start":
		msg.ReplyMarkup = mainMenu
		mainMenu.ResizeKeyboard = true
		updateState(message.Chat.ID, "main_menu")
	case "Описать проблему":
		msg.Text = "Выберите тип проблемы:"
		msg.ReplyMarkup = describeProblemMenu
		describeProblemMenu.ResizeKeyboard = true
		updateState(message.Chat.ID, "describe_problem")
	case "Вернуться в предыдущее меню":
		handleReturnToPreviousMenu(message.Chat.ID)
		return
	default:
		// Обработка других сообщений в зависимости от текущего состояния
		handleOtherMessages(message)
	}

	bot.Send(msg)
}

func handleOtherMessages(message *tgbotapi.Message) {
	userState := getUserState(message.Chat.ID)
	switch userState.CurrentState {
	case "main_menu":
		// Обработка сообщений в основном меню
	case "describe_problem":
		// Обработка сообщений в разделе "Описать проблему"
		// Можно добавить дополнительную логику в зависимости от выбора пользователя
	default:
		// Обработка сообщений в других состояниях
	}
}

func handleReturnToPreviousMenu(chatID int64) {
	userState := getUserState(chatID)
	if userState.PrevState != "" {
		// Вернуться в предыдущее меню
		msg := tgbotapi.NewMessage(chatID, "Вы вернулись в предыдущее меню.")
		msg.ReplyMarkup = mainMenu
		mainMenu.ResizeKeyboard = true
		updateState(chatID, userState.PrevState)
		bot.Send(msg)
	} else {
		// Если предыдущего меню нет, отправить информацию пользователю
		msg := tgbotapi.NewMessage(chatID, "Нет предыдущего меню.")
		bot.Send(msg)
	}
}

// Вспомогательные функции для работы с состояниями
func updateState(chatID int64, newState string) {
	userStatesMutex.Lock()
	defer userStatesMutex.Unlock()

	userState, ok := userStates[chatID]
	if !ok {
		userState = &UserState{}
		userStates[chatID] = userState
	}

	userState.PrevState = userState.CurrentState
	userState.CurrentState = newState
}

func getUserState(chatID int64) *UserState {
	userStatesMutex.Lock()
	defer userStatesMutex.Unlock()

	userState, ok := userStates[chatID]
	if !ok {
		userState = &UserState{}
		userStates[chatID] = userState
	}

	return userState
}
