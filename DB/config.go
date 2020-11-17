package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

var DB *sql.DB
var once sync.Once

// New to create a DB instance for one time
func New() *sql.DB {

	once.Do(func() {
		logger, _ := thoth.Init("log")

		user, exist := os.LookupEnv("DB_USER")

		if !exist {
			logger.Log(errors.New("DB_USER not set in .env"))
			log.Fatal("DB_USER not set in .env")
		}

		pass, exist := os.LookupEnv("DB_PASS")

		if !exist {
			logger.Log(errors.New("DB_PASS not set in .env"))
			log.Fatal("DB_PASS not set in .env")
		}

		host, exist := os.LookupEnv("DB_HOST")

		if !exist {
			logger.Log(errors.New("DB_HOST not set in .env"))
			log.Fatal("DB_HOST not set in .env")
		}

		credentials := fmt.Sprintf("%s:%s@tcp(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)

		database, err := sql.Open("mysql", credentials)

		if err != nil {
			logger.Log(err)
			log.Fatal(err)
		} else {
			fmt.Println("Database Connection Successful")
		}
		/*_, err = database.Exec(`DROP DATABASE Promotions;`)

		if err != nil {
			fmt.Println(err)
		}*/

		_, err = database.Exec(`CREATE DATABASE Promotions;`)

		if err != nil {
			fmt.Println(err)
		}

		_, err = database.Exec(`USE Promotions;`)

		if err != nil {
			fmt.Println(err)
		}

		_, err = database.Exec(`
		CREATE TABLE UserPromo (
			promo_id int NOT NULL AUTO_INCREMENT,
			title varchar(45) NOT NULL,
			description varchar(100) NOT NULL,
			start_date varchar(45) NOT NULL,
			end_date varchar(45) NOT NULL,
			PRIMARY KEY (promo_id)
		  );
	  `)

		if err != nil {
			fmt.Println(err)
		}

		_, err = database.Exec(`
		CREATE TABLE INCPromo (
			promo_id int NOT NULL AUTO_INCREMENT,
			title varchar(45) NOT NULL,
			description varchar(100) NOT NULL,
			start_date varchar(45) NOT NULL,
			end_date varchar(45) NOT NULL,
			PRIMARY KEY (promo_id)
		  );
	  `)

		if err != nil {
			fmt.Println(err)
		}
		DB = database
	})
	return DB
}
