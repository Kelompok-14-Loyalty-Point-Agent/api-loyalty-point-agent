package mysql_driver

import (
	"api-loyalty-point-agent/drivers/mysql/providers"
	"api-loyalty-point-agent/drivers/mysql/stock_details"
	"api-loyalty-point-agent/drivers/mysql/stock_transactions"
	"api-loyalty-point-agent/drivers/mysql/stocks"
	"api-loyalty-point-agent/drivers/mysql/users"
	"os"
	"time"

	aws_driver "api-loyalty-point-agent/drivers/aws"

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
	err := db.AutoMigrate(&users.User{}, &providers.Provider{}, &stocks.Stock{}, &stock_details.StockDetail{}, &stock_transactions.StockTransaction{})

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
	} else {
		result := db.Create(&admin)

		if result.Error != nil {
			log.Fatalf("failed to create admin: %s\n", result.Error)
		}

		log.Println("admin created")
	}
}

func SeedProvider(db *gorm.DB) {
	var provider1 = providers.Provider{
		Name: "Telkomsel",
		URL:  "",
	}

	var provider2 = providers.Provider{
		Name: "XL",
		URL:  "",
	}

	var record providers.Provider

	_ = db.First(&record, "name = ?", provider1.Name)

	if record.ID != 0 {
		log.Printf("provider already exists\n")
	} else {
		filePath := "./assets/providers/tsel.png"

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		defer file.Close()

		url, err := aws_driver.UploadFileToBucket("tsel.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider1.URL = url

		result := db.Create(&provider1)

		if result.Error != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		filePath = "./assets/providers/xl.png"

		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}
		defer file.Close()

		url, err = aws_driver.UploadFileToBucket("xl.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider2.URL = url

		result = db.Create(&provider2)

		if result.Error != nil {
			log.Fatalf("failed to create admin: %s\n", err.Error())
		}

		log.Println("2 providers created")
	}

}

func SeedStock(db *gorm.DB) {
	var stock1 = stocks.Stock{
		Type:       "data",
		TotalStock: 500,
		ProviderID: 1,
		LastTopUp: time.Now(),
	}

	var stock2 = stocks.Stock{
		Type:       "credit",
		TotalStock: 1000000,
		ProviderID: 1,
		LastTopUp: time.Now(),
	}

	var stock3 = stocks.Stock{
		Type:       "data",
		TotalStock: 500,
		ProviderID: 2,
		LastTopUp: time.Now(),
	}

	var stock4 = stocks.Stock{
		Type:       "credit",
		TotalStock: 1000000,
		ProviderID: 2,
		LastTopUp: time.Now(),
	}

	var record stocks.Stock

	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("stock already exists\n")
	} else {
		result := db.Create(&stock1)

		if result.Error != nil {
			log.Fatalf("failed to create stock1: %s\n", result.Error)
		}

		result = db.Create(&stock2)

		if result.Error != nil {
			log.Fatalf("failed to create stock2: %s\n", result.Error)
		}

		result = db.Create(&stock3)

		if result.Error != nil {
			log.Fatalf("failed to create stock3: %s\n", result.Error)
		}

		result = db.Create(&stock4)

		if result.Error != nil {
			log.Fatalf("failed to create stock4: %s\n", result.Error)
		}

		log.Println("4 stocks created")
	}
}
