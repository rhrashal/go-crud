package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Option A: Use full DB_URL from .env (as in blog)
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// Option B: Build it from separate env vars (more control)
		host := os.Getenv("DB_HOST") // default: localhost
		user := os.Getenv("DB_USER") // default: postgres
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME") // ← your custom name here
		port := os.Getenv("DB_PORT")   // default: 5432

		if host == "" {
			host = "localhost"
		}
		if user == "" {
			user = "postgres"
		}
		if password == "" {
			password = "abcd"
		}
		if dbname == "" {
			dbname = "myappdb"
		} // fallback
		if port == "" {
			port = "5432"
		}

		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
			host, user, password, dbname, port)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")
}
