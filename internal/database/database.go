package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "postgres@localhost"
)

type User struct {
	ID    int
	Token string
}

type Answer struct {
	ID         int
	Expression string
	Answer     float64
}

func main() {
	// Строка подключения к базе данных PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Соединение с базой данных установлено")

	// Создание таблицы пользователей (users)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		token TEXT
	);`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Таблица пользователей создана успешно")

	// Создание таблицы ответов (answers)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS answers (
		id SERIAL PRIMARY KEY,
		expression TEXT,
		answer REAL
	);`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Таблица ответов создана успешно")

	// Пример вставки данных в таблицу пользователей
	_, err = db.Exec("INSERT INTO users (token) VALUES ($1)", "example_token")
	if err != nil {
		panic(err)
	}

	fmt.Println("Данные успешно добавлены в таблицу пользователей")

	// Пример вставки данных в таблицу ответов
	_, err = db.Exec("INSERT INTO answers (expression, answer) VALUES ($1, $2)", "2 + 2", 4.0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Данные успешно добавлены в таблицу ответов")
}
