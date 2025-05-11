package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"psy_match/config"
	"time"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

// UserProfile – структура данных пользователя
type UserProfile struct {
	UserID           int64  `json:"user_id"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	BirthDate        string `json:"dob"`
	ZodiacSign       string `json:"zodiac_sign"`
	MBTI             string `json:"mbti"`
	TelegramUsername string `json:"telegram_username"`
}

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.AppConfig.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}
	if conn == nil {
		return nil, fmt.Errorf("Ошибка: соединение с базой данных не установлено")
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}
	// Применяем миграции
	err = applyMigration(conn)
	if err != nil {
		return nil, err
	}
	db = conn
	log.Println("Подключение к базе данных установлено.")
	return conn, nil
}

func applyMigration(conn *pgx.Conn) error {
	migrationSQL, err := ioutil.ReadFile("./migrations/001_init.up.sql")
	if err != nil {
		return fmt.Errorf("Ошибка при чтении миграции: %v", err)
	}
	_, err = conn.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		return fmt.Errorf("Ошибка при выполнении миграции: %v", err)
	}
	log.Println("Миграция выполнена успешно.")
	return nil
}

func GetZodiacSign(birthDate string) string {
	layout := "2006-01-02"
	date, err := time.Parse(layout, birthDate)
	if err != nil {
		log.Printf("Ошибка парсинга даты: %v", err)
		return ""
	}

	month := date.Month()
	day := date.Day()

	switch {
	case (month == 3 && day >= 21) || (month == 4 && day <= 19):
		return "Овен"
	case (month == 4 && day >= 20) || (month == 5 && day <= 20):
		return "Телец"
	case (month == 5 && day >= 21) || (month == 6 && day <= 20):
		return "Близнецы"
	case (month == 6 && day >= 21) || (month == 7 && day <= 22):
		return "Рак"
	case (month == 7 && day >= 23) || (month == 8 && day <= 22):
		return "Лев"
	case (month == 8 && day >= 23) || (month == 9 && day <= 22):
		return "Дева"
	case (month == 9 && day >= 23) || (month == 10 && day <= 22):
		return "Весы"
	case (month == 10 && day >= 23) || (month == 11 && day <= 21):
		return "Скорпион"
	case (month == 11 && day >= 22) || (month == 12 && day <= 21):
		return "Стрелец"
	case (month == 12 && day >= 22) || (month == 1 && day <= 19):
		return "Козерог"
	case (month == 1 && day >= 20) || (month == 2 && day <= 18):
		return "Водолей"
	case (month == 2 && day >= 19) || (month == 3 && day <= 20):
		return "Рыбы"
	default:
		return ""
	}
}

func CalculateAge(birthDate string) int {
	layout := "2006-01-02"
	date, err := time.Parse(layout, birthDate)
	if err != nil {
		log.Printf("Ошибка парсинга даты: %v", err)
		return 0
	}
	now := time.Now()
	age := now.Year() - date.Year()
	if now.YearDay() < date.YearDay() {
		age--
	}
	return age
}

// SaveUserProfile – Сохранение/обновление профиля в БД
func SaveUserProfile(userId int64, data map[string]string) {

	// Определяем знак зодиака (возраст не вычисляем при сохранении)
	birthDate := data["dob"]
	zodiac := GetZodiacSign(birthDate)

	query := `INSERT INTO user_profiles (user_id, name, birth_date, zodiac_sign, telegram_username)
	          VALUES ($1, $2, $3, $4, $5)
	          ON CONFLICT (user_id) DO UPDATE SET 
	            name = EXCLUDED.name, birth_date = EXCLUDED.birth_date, 
	            zodiac_sign = EXCLUDED.zodiac_sign,
	            telegram_username = EXCLUDED.telegram_username`

	_, err := db.Exec(context.Background(), query, userId, data["name"], birthDate, zodiac, data["telegram_username"])
	if err != nil {
		log.Fatalf("Ошибка при сохранении профиля: %v", err)
	}
	log.Println("Профиль сохранён.")
}

func SaveTestResult(userID int64, mbti string) error {
	query := `INSERT INTO mbti_results (user_id, mbti)
	          VALUES ($1, $2)
	          ON CONFLICT (user_id) DO UPDATE SET mbti = EXCLUDED.mbti`

	_, err := db.Exec(context.Background(), query, userID, mbti)
	if err != nil {
		log.Printf("Ошибка при сохранении htpультатов теста: %v", err)
		return err
	}

	log.Printf("Результат теста сохранён для пользователя %d: %s", userID, mbti)
	return nil
}

func GetUserProfileByID(userID int64) (*UserProfile, int, error) {
	if db == nil {
		return nil, 0, fmt.Errorf("Соединение с БД не установлено")
	}

	log.Printf("Поиск профиля по user_id: %d", userID)

	query := `SELECT u.user_id, u.name, u.birth_date, u.zodiac_sign, u.telegram_username, m.mbti
	          FROM user_profiles u
	          LEFT JOIN mbti_results m ON u.user_id = m.user_id
	          WHERE u.user_id = $1`

	row := db.QueryRow(context.Background(), query, userID)
	var profile UserProfile
	err := row.Scan(&profile.UserID, &profile.Name, &profile.BirthDate, &profile.ZodiacSign, &profile.TelegramUsername, &profile.MBTI)
	if err != nil {
		return nil, 0, err
	}

	age := CalculateAge(profile.BirthDate)
	return &profile, age, nil
}

func GetUserProfileByTelegram(username string) (*UserProfile, int, error) {
	if db == nil {
		return nil, 0, fmt.Errorf("Соединение с БД не установлено")
	}

	query := `SELECT u.user_id, u.name, u.birth_date, u.zodiac_sign, u.telegram_username, m.mbti
	          FROM user_profiles u
	          LEFT JOIN mbti_results m ON u.user_id = m.user_id
	          WHERE u.telegram_username = $1`
	row := db.QueryRow(context.Background(), query, username)

	var profile UserProfile
	err := row.Scan(&profile.UserID, &profile.Name, &profile.BirthDate, &profile.ZodiacSign, &profile.TelegramUsername, &profile.MBTI)
	if err != nil {
		return nil, 0, err
	}

	age := CalculateAge(profile.BirthDate)
	return &profile, age, nil
}
