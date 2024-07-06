package texts

import "fmt"

func GreetingText(name string) string {
	return fmt.Sprintf("Привет, %s! Я KIWI BOT 😍 \nУже миллионы людей знакомятся в KIWI 🥝\nЯ помогу найти тебе пару или просто друзей 👫", name)
}

const (
	AgeComplete  = "Ваш возраст сохранен 🎉"
	AgeQuestion  = "Сколько тебе лет? 🎂"
	AgeLower     = "Ты должен быть старше 16-ти лет! 🚫"
	AgeUpper     = "Пожалуйста, введите Ваш реальный возраст 📅"
	AgeIncorrect = "Укажите ваш реальный возраст одним числом 🔢"
)

const (
	GenderQuestion = "Кто вы? 👫"
	GenderMale     = "Я парень 👦"
	GenderFemale   = "Я девушка 👧"
)

const (
	PhotoInfo         = "Пожалуйста, отправьте мне свою фотографию 📷\nИли вы можете оставить ту, которая у вас уже стоит в Telegram 😊"
	PhotoDefault      = "Оставить фотографию из Telegram 😊"
	PhotoDefaultEmpty = "Не удалось получить ни одну фотографию, загрузите пожалуйста свою настоящюю фотографию 📸"
	PhotoEmpty        = "Отправьте пожалуйста свою настоящюю фотографию 📸"
	PhotoComplete     = "Ваша фотография сохранена 🎉"
)

const (
	AboutQuestion = "Расскажите немного о себе 📝"
	AboutComplete = "Информация сохранена 🎉"
)

func AboutLongText(num int, max int) string {
	return fmt.Sprintf("Слишком длинный текст 🤔\n%d/%d", num, max)
}

const (
	ProfileFillInfo = "Сейчас я задам тебе несколько вопросов, чтобы получить минимальную информацию для начала поиска 📋\nТвоя задача — честно на них ответить 🥝\nЕсли допустишь ошибку в ответе, не волнуйся! Ты всегда сможешь вернуться и изменить ответ 🔄\n\n🌟 Давай начнем! 🌟"
)

const (
	LocationQuestion = "Отправьте мне своё местоположение 🗺️\nИли напишите город в котором проживаете"
	LocationSend     = "Отправить местоположение 🗺️"
	LocationComplete = "Местоположение сохранено 🎉"
)
