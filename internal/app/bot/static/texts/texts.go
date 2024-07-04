package texts

import "fmt"

func GreetingText(name string) string {
	return fmt.Sprintf("Привет, %s! Я KIWI BOT 😍 \nУже миллионы людей знакомятся в KIWI 🥝\nЯ помогу найти тебе пару или просто друзей 👫", name)
}

const (
	AgeQuestion  = "Сколько Вам лет?"
	AgeLower     = "Вы должны быть старше 16-ти лет!"
	AgeUpper     = "Пожалуйста, введите ваш реальный возраст"
	AgeIncorrect = "Укажите ваш реальный возраст одним числом"
)

const (
	GenderQuestion = "Кто вы?"
	GenderMale     = "Я парень"
	GenderFemale   = "Я девушка"
)
