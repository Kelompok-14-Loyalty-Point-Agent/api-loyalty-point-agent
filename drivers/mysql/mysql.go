package mysql_driver

import (
	"api-loyalty-point-agent/drivers/mysql/users"

	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

// connect to the database
func (config *DBConfig) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")

	return db
}

// perform migration
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&users.User{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func SeedAdmin(db *gorm.DB) {
	var admin = users.User{
		Name:     "admin",
		Password: "admin123",
		Email:    "admin@example.com",
		Role:     "admin",
	}

	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("failed to hash admin password: %s\n", err)
	}

	admin.Password = string(password)
	
	var record users.User

	_ = db.First(&record, "email = ?", admin.Email)

	if record.ID != 0 {
		log.Printf("admin already exists\n")
	}else{
		result := db.Create(&admin)

		if result.Error != nil {
			log.Fatalf("failed to create admin: %s\n", result.Error)
		}

		log.Println("admin created")
	}
}
