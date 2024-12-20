package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql" 
)

var db *sql.DB

// Функция для установки и инициализации БД
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблиц по отдельности
	tables := []string{
		`CREATE TABLE IF NOT EXISTS admins (
			id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			username VARCHAR(200) NOT NULL,
			password VARCHAR(50) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS employees (
			id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(200) NOT NULL,
			password VARCHAR(50) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS clients (
			id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(200) NOT NULL,
			birth_date varchar(15) NOT NULL,
			phone_number VARCHAR(20) NOT NULL,
			adres VARCHAR(100) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS subscriptions (
			id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL,
			price DECIMAL(10,2) NOT NULL,
			sold_date DATE NOT NULL,
			client_id INT NOT NULL, 
			employee_id INT NOT NULL,
			FOREIGN KEY (client_id) REFERENCES clients(id),
			FOREIGN KEY (employee_id) REFERENCES employees(id)
		);`,
	}

	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Получаем доступ к БД
func GetDB() *sql.DB {
	return db
}