package database

// импортируем нужное
import (
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db *sql.BD

// фукнция для установки инициализации бд
func initDB(dataSourceName string) {
	var err error
	db,err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	create table if not exists admins (
		id int primary key auto_increment,
		username varchar(200) not null,
		password varchar(50) not null
	);
	
	create tabe if not exists employees (
		id int primary key auto_increment,
		name varchar(200) not null,
		password varhcar(50) not null
	);

	create table if not exists clients (
		id int primary key auto_increment,
		name varchar(200) not null,
		birtch_date date not null,
		phone_number varchar(20) not null,
		adres varchar(100) not null
	);

	create table if not exists subscriptions (
		id int primary key auto_increment,
		name varchar(100) not null,
		price decimal(10,2) not null,
		sold_date date not null,
		client_id int not null, 
		employee_id int not null,
		foreign key (client_id) references clients(id),
		foreign key (employee_id) references employees(id)
	); 
	`)

	if err != nil {
		log.Fatal(err)
	}
} 

// получаем доступ к бд
func GetDB() *sql.DB {
	return db
}