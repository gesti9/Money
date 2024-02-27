package main

// import (
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("6796961656:AAGimXMVJzd0a1JwkFvSEqR28mbMQr2aL1k")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message == nil { // ignore any non-Message Updates
// 			continue
// 		}

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Привет:)
// Напиши нам о своей проблеме!`)
// 		keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 			tgbotapi.NewInlineKeyboardRow(
// 				tgbotapi.NewInlineKeyboardButtonURL("Посетите мой сайт", "https://t.me/Bernar25"),
// 			),
// 		)
// 		msg.ReplyMarkup = keyboard

// 		bot.Send(msg)
// 	}
// }

import (
	"fmt"
	"log"
	"money/server"
	"os"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botToken               = "6796961656:AAGimXMVJzd0a1JwkFvSEqR28mbMQr2aL1k"
	CommandSendApplication = "SEND_APPLICATION"
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
			tgbotapi.NewKeyboardButton("Наши услуги!"),
			tgbotapi.NewKeyboardButton("Ремонт и выкуп!"),
		),
	)
)

func main() {
	// User("Напишите свой номер сударь") //Отправка пользаку письмо
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			res := update.Message.Text
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)

			switch update.Message.Text {
			case "/start":
				Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, можешь описать свою проблему и не забудь оставить свои контакты(WhatsApp, Telegram) мы с тобой свяжемся! с/у KazSync:)")
				msg.ReplyMarkup = mainMenu
				bot.Send(msg)
			case "Ремонт и выкуп!":
				Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `	Привет:)
				Высылай все сюда👇🤗`)
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Сontact service", "https://t.me/era_gl450"),
					),
					tgbotapi.NewInlineKeyboardRow(

						tgbotapi.NewInlineKeyboardButtonURL("Сontact service WhatsApp", "https://wa.me/77473195507"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			case "Наши услуги!":
				Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `🖥️🛠️ Наши услуги 🛠️🖥️
				🤝 Простое решение компьютерных проблем
				Помогаем тем, кто не разбирается в компьютерах:
				
				🎓 Обучение и консультации: Индивидуальные уроки и консультации по компьютерной грамотности.
				
				🔧 Техническая поддержка: Устранение ошибок, настройка ПО, восстановление данных.
				
				🚀 Удаленное обслуживание: Онлайн-поддержка и планы обслуживания для оптимальной производительности.
				
				🔍 Пользовательские решения: Настройка почты, социальных сетей, помощь с программами.
				
				🔄 Обновление и советы: Информация о последних технологиях и рекомендации по обновлению.
				
				🚑 Срочный выкуп и ремонт
				💸 Выкуп: Предоставляем услуги выкупа техники по справедливой цене.
				
				🔨 Ремонт: Экспресс-ремонт компьютеров, ноутбуков и гаджетов.
				
				Доверьте нам заботу о вашей технике — от решения проблем до выкупа и ремонта.`)
				bot.Send(msg)
			default:
				Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				Mess("ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "\n" + update.Message.Text)
			}

		}
	}
}

func Mess(text string) {
	c := server.New(botToken)

	c.SendMessage(text, -1002022179282)
}
func User(text string) {
	c := server.New(botToken)

	c.SendMessage(text, 308722033)
	fmt.Println("Отправил!")
}

func Log(n string) {
	// Получаем текущую дату и время
	currentTime := time.Now()

	// Форматируем дату и время в строку
	dateTimeString := currentTime.Format("2006-01-02 15:04:05")

	// Имя файла, в который мы будем записывать данные
	fileName := "log.txt"

	// Открываем файл для записи (если файла нет, он будет создан)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Записываем дату и время в файл
	_, err = file.WriteString(dateTimeString + "    " + n)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}
