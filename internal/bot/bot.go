package bot

import (
	"fmt"
	telebot "gopkg.in/telebot.v3"
	"log"
	"psy_match/config"
	"psy_match/internal/database"
	"psy_match/internal/services"
	"strconv"
	"strings"
	"time"
)

var userStates = make(map[int64]*UserState)
var userTestStates = make(map[int64]*TestState)

type UserState struct {
	Step   string
	UserID int64
	Data   map[string]string
}

type TestState struct {
	UserID  int64
	Answers []string
	Step    int
}

func StartBot() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.AppConfig.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10},
	})
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Handle("/start", StartHandler)
	bot.Handle("/profile", ProfileHandler)
	bot.Handle("/test", TestMenuHandler)
	bot.Handle("/match", CompatibilityHandler)
	bot.Handle(telebot.OnCallback, func(c telebot.Context) error {
		log.Printf("Получены данные callback: [%s]", c.Data())
		data := strings.TrimSpace(c.Data())
		switch data {
		case "start_test":
			return StartTestHandler(c)
		case "start_profile":
			return StartProfileHandler(c)
		case "I", "E", "T", "F", "S", "N", "J", "P":
			return handleTestAnswer(c)
		default:
			log.Println("Неизвестный callback:", data)
			return c.Send("Ошибка: неизвестный запрос.")
		}
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		log.Println("Получено сообщение:", c.Text())
		return TextHandler(c)
	})

	log.Println("Бот запущен...")
	bot.Start()
}

func StartHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	profile, age, err := database.GetUserProfileByID(userID)
	if err != nil || profile == nil {
		userStates[userID] = &UserState{UserID: userID, Step: "name", Data: make(map[string]string)}
		return c.Send("Привет! Давайте начнем с заполнения Вашего профиля.\nВведите ваше имя:")
	}
	return c.Send(fmt.Sprintf("Ваш профиль:\nИмя: %s\nВозраст: %d\nЗнак зодиака: %s", profile.Name, age, profile.ZodiacSign))
}

func ProfileHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	profile, age, err := database.GetUserProfileByID(userID)

	inlineKeys := &telebot.ReplyMarkup{}
	if err != nil || profile == nil {
		btnFillProfile := inlineKeys.Data("Заполнить профиль", "start_profile")
		inlineKeys.Inline(inlineKeys.Row(btnFillProfile))
		return c.Send("Вы ещё не заполнили профиль.\n Хотите заполнить его?", inlineKeys)
	}

	message := fmt.Sprintf("Имя: %s\nВозраст: %d\nЗнак зодиака: %s",
		profile.Name, age, profile.ZodiacSign)

	btnFillProfile := inlineKeys.Data("Заполнить профиль", "start_profile")
	inlineKeys.Inline(inlineKeys.Row(btnFillProfile))

	return c.Send(message, inlineKeys)
}

func StartProfileHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)

	if userID == 0 {
		return c.Send("Ошибка: не удалось получить ваш ID.")
	}

	userStates[userID] = &UserState{
		UserID: userID,
		Step:   "name",
		Data:   make(map[string]string),
	}
	return c.Send("Введите ваше имя:")
}

func ProfileStepHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	state := userStates[userID]

	switch state.Step {
	case "name":
		state.Data["name"] = c.Text()
		state.Step = "dob"
		return c.Send("Введите вашу дату рождения.\nВведите год:")

	case "dob":
		text := c.Text()
		year, err := strconv.Atoi(text)
		if err != nil || year < 1900 || year > time.Now().Year() {
			return c.Send("Введите корректный год (например, 1995):")
		}

		state.Data["year"] = text
		state.Step = "dob_month"

		return c.Send("Введите месяц рождения (числом, например, 07):")

	case "dob_month":
		text := c.Text()
		month, err := strconv.Atoi(text)
		if err != nil || month < 1 || month > 12 {
			return c.Send("Введите корректный месяц (01-12):")
		}
		state.Data["month"] = fmt.Sprintf("%02d", month)
		state.Step = "dob_day"
		return c.Send("Введите день рождения (например, 28):")

	case "dob_day":
		text := c.Text()
		day, err := strconv.Atoi(text)
		if err != nil || day < 1 || day > 31 {
			return c.Send("Введите корректный день (01-31):")
		}

		year := userStates[userID].Data["year"]
		month := userStates[userID].Data["month"]
		dateStr := fmt.Sprintf("%s-%s-%02d", year, month, day)

		_, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.Send("Введена некорректная дата.\n Попробуйте ещё раз:")
		}

		state.Data["dob"] = dateStr

		telegramUsername := c.Sender().Username
		if telegramUsername != "" {
			formattedUsername := fmt.Sprintf("@%s", telegramUsername)
			state.Data["telegram_username"] = formattedUsername
		} else {
			state.Data["telegram_username"] = "unknown"
		}

		database.SaveUserProfile(state.UserID, state.Data)

		delete(userStates, userID)

		profile, _, err := database.GetUserProfileByID(userID)
		if err != nil || profile == nil {
			return c.Send("Произошла ошибка при получении профиля.")
		}

		inlineKeys := &telebot.ReplyMarkup{}
		if profile.MBTI == "" {
			btnStartTest := inlineKeys.Data("Пройти тест", "start_test")
			inlineKeys.Inline(inlineKeys.Row(btnStartTest))
			return c.Send("Профиль сохранен!", inlineKeys)
		} else {
			return c.Send("Профиль сохранен!")
		}
	}
	return nil
}

func TestMenuHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	profile, _, err := database.GetUserProfileByID(userID)

	inlineKeys := &telebot.ReplyMarkup{}

	if profile != nil && profile.MBTI == "" {
		btnStartTest := inlineKeys.Data("Пройти тест", "start_test")
		inlineKeys.Inline(inlineKeys.Row(btnStartTest))
		return c.Send("Вы ещё не проходили тест.\n Хотите пройти?", inlineKeys)
	}

	if err != nil || profile == nil {
		btnStartTest := inlineKeys.Data("Заполнить профиль", "start_profile")
		inlineKeys.Inline(inlineKeys.Row(btnStartTest))
		return c.Send("Вы ещё не зарегистрировались.\nПерейдите к заполнению профиля", inlineKeys)
	}

	desc := services.GetPersonalitySelfDescription(profile.MBTI)
	message := fmt.Sprintf("Ваш результат: %s", desc)

	btnRetakeTest := inlineKeys.Data("Пройти тест", "start_test")
	inlineKeys.Inline(inlineKeys.Row(btnRetakeTest))

	return c.Send(message, inlineKeys)
}

func StartTestHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)

	userTestStates[userID] = &TestState{UserID: userID, Step: 1, Answers: []string{}}

	return sendTestQuestion(c, c.Sender(), userTestStates[userID])
}

func sendTestQuestion(c telebot.Context, recipient *telebot.User, state *TestState) error {
	var question string
	var inlineButtons [][]telebot.InlineButton

	switch state.Step {
	case 1:
		question = "1. Как вы предпочитаете проводить свободное время?"
		inlineButtons = [][]telebot.InlineButton{
			{
				{Text: "В одиночестве", Data: "I"},
				{Text: "С друзьями", Data: "E"},
			},
		}
	case 2:
		question = "2. Как вы принимаете решения?"
		inlineButtons = [][]telebot.InlineButton{
			{
				{Text: "Анализирую факты", Data: "S"},
				{Text: "Доверяюсь чувствам", Data: "N"},
			},
		}
	case 3:
		question = "3. В конфликтной ситуации вы обычно:"
		inlineButtons = [][]telebot.InlineButton{
			{
				{Text: "Избегаю конфликта", Data: "F"},
				{Text: "Обсуждаю открыто", Data: "T"},
			},
		}
	case 4:
		question = "4. Ваш выбор в спорной ситуации?"
		inlineButtons = [][]telebot.InlineButton{
			{
				{Text: "Рационально", Data: "J"},
				{Text: "Эмоционально", Data: "P"},
			},
		}
	}
	return c.Send(question, &telebot.ReplyMarkup{
		InlineKeyboard: inlineButtons,
	})
}

func handleTestAnswer(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	answer := c.Data()

	state, exists := userTestStates[userID]
	if !exists {
		return c.Send("Вы не начали тест.")
	}

	state.Answers = append(state.Answers, answer)

	state.Step++

	if state.Step > 4 {
		mbti := strings.Join(state.Answers, "")

		err := database.SaveTestResult(state.UserID, mbti)
		if err != nil {
			return c.Send("Произошла ошибка при сохранении результата.")
		}

		desc := services.GetPersonalitySelfDescription(mbti)

		err = c.Send(fmt.Sprintf("Тест завершен! Вы: %s", desc))
		if err != nil {
			return err
		}

		err = c.Send("Введите @telegram_username пользователя для расчета совместимости:")
		if err != nil {
			return err
		}

		userStates[state.UserID] = &UserState{UserID: state.UserID, Step: "nickname", Data: make(map[string]string)}

		return nil
	}

	return sendTestQuestion(c, c.Sender(), state)
}

func TextHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	if state, exists := userStates[userID]; exists {
		if state.Step == "nickname" {
			return HandleCompatibilityAnswer(c)
		}
		return ProfileStepHandler(c)
	}
	return c.Send("Я не понимаю этот запрос.\nИспользуйте команды для работы со мной в меню слева.")
}

func CompatibilityHandler(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	c.Send("Введите @telegram_username пользователя для расчета совместимости:")

	userStates[userID] = &UserState{UserID: userID, Step: "nickname", Data: make(map[string]string)}
	return nil
}

func HandleCompatibilityAnswer(c telebot.Context) error {
	userID := int64(c.Sender().ID)
	otherUsername := c.Text()

	profile, _, err := database.GetUserProfileByID(userID)
	if err != nil || profile == nil {
		return c.Send("Не удалось найти ваш профиль.")
	}
	otherProfile, _, err := database.GetUserProfileByTelegram(otherUsername)
	if err != nil || otherProfile == nil {
		return c.Send("Не удалось найти профиль другого пользователя.")
	}

	mbtiCompatibility, _ := services.GetCompatibilityData(profile.MBTI, otherProfile.MBTI)

	zodiacCompatibility := services.GetZodiacCompatibility(profile.ZodiacSign, otherProfile.ZodiacSign)
	fmt.Println(mbtiCompatibility)
	fmt.Println(zodiacCompatibility)
	compatibilityPerc := mbtiCompatibility + zodiacCompatibility
	fmt.Println(compatibilityPerc)

	return c.Send(fmt.Sprintf("Совместимость между Вами и %s: %d%%\n", otherUsername, compatibilityPerc))
}
