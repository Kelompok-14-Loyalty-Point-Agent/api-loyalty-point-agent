package mysql_driver

import (
	"api-loyalty-point-agent/drivers/mysql/providers"
	"api-loyalty-point-agent/drivers/mysql/stock_details"
	"api-loyalty-point-agent/drivers/mysql/stock_transactions"
	"api-loyalty-point-agent/drivers/mysql/stocks"
	"api-loyalty-point-agent/drivers/mysql/transactions"
	"api-loyalty-point-agent/drivers/mysql/users"
	"api-loyalty-point-agent/drivers/mysql/profiles"

	"os"

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
	err := db.AutoMigrate(&profiles.Profile{}, &users.User{}, &providers.Provider{}, &stocks.Stock{}, &stock_details.StockDetail{}, &stock_transactions.StockTransaction{}, &transactions.Transaction{})

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
		ProfileID: 0,
		Profile: profiles.Profile{},
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
		filePath := "./assets/users/user.png"

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create admin: %s\n", err.Error())
		}

		defer file.Close()

		url, err := aws_driver.UploadFileToBucket("user.png", file)
		if err != nil {
			log.Fatalf("failed to create admin: %s\n", err.Error())
		}

		err = aws_driver.DownloadFileFromBucket(url, "./assets/users/")
		if err != nil {
			log.Fatalf("failed to create admin: %s\n", err.Error())
		}

		var profile profiles.Profile

		profile.URL = url

		result := db.Create(&profile)

		if result.Error != nil {
			log.Fatalf("failed to create admin's profile: %s\n", result.Error)
		}

		admin.ProfileID = profile.ID
		admin.Profile = profile

		result = db.Create(&admin)

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

	var provider3 = providers.Provider{
		Name: "Smartfren",
		URL:  "",
	}

	var provider4 = providers.Provider{
		Name: "Indosat Ooredoo",
		URL:  "",
	}

	var provider5 = providers.Provider{
		Name: "Axis",
		URL:  "",
	}

	var provider6 = providers.Provider{
		Name: "3",
		URL:  "",
	}

	var record providers.Provider

	_ = db.First(&record, "name = ?", provider1.Name)

	if record.ID != 0 {
		log.Printf("provider already exists\n")
	} else {
		// tsel
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

		// xl
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
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		// smartfren
		filePath = "./assets/providers/smartfren.png"

		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}
		defer file.Close()

		url, err = aws_driver.UploadFileToBucket("smartfren.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider3.URL = url

		result = db.Create(&provider3)

		if result.Error != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		// indosat ooredoo
		filePath = "./assets/providers/indosat.png"

		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}
		defer file.Close()

		url, err = aws_driver.UploadFileToBucket("indosat.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider4.URL = url

		result = db.Create(&provider4)

		if result.Error != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		// axis
		filePath = "./assets/providers/axis.png"

		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}
		defer file.Close()

		url, err = aws_driver.UploadFileToBucket("axis.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider5.URL = url

		result = db.Create(&provider5)

		if result.Error != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		// 3
		filePath = "./assets/providers/3.png"

		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}
		defer file.Close()

		url, err = aws_driver.UploadFileToBucket("3.png", file)
		if err != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		provider6.URL = url

		result = db.Create(&provider6)

		if result.Error != nil {
			log.Fatalf("failed to create provider: %s\n", err.Error())
		}

		log.Println("6 providers created")
	}

}

func SeedStock(db *gorm.DB) {
	// tsel stock
	var stock1 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 1,
		// LastTopUp:  time.Now(),
	}

	var stock2 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 1,
		// LastTopUp:  time.Now(),
	}

	// xl stock
	var stock3 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 2,
		// LastTopUp:  time.Now(),
	}

	var stock4 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 2,
		// LastTopUp:  time.Now(),
	}

	// smartfren stock
	var stock5 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 3,
	}

	var stock6 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 3,
	}

	// indosat stock
	var stock7 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 4,
	}

	var stock8 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 4,
	}

	// axis stock
	var stock9 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 5,
	}

	var stock10 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 5,
	}

	// 3 stock
	var stock11 = stocks.Stock{
		Type:       "data",
		TotalStock: 0,
		ProviderID: 6,
	}

	var stock12 = stocks.Stock{
		Type:       "credit",
		TotalStock: 0,
		ProviderID: 6,
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

		result = db.Create(&stock5)

		if result.Error != nil {
			log.Fatalf("failed to create stock5: %s\n", result.Error)
		}

		result = db.Create(&stock6)

		if result.Error != nil {
			log.Fatalf("failed to create stock6: %s\n", result.Error)
		}

		result = db.Create(&stock7)

		if result.Error != nil {
			log.Fatalf("failed to create stock7: %s\n", result.Error)
		}

		result = db.Create(&stock8)

		if result.Error != nil {
			log.Fatalf("failed to create stock8: %s\n", result.Error)
		}

		result = db.Create(&stock9)

		if result.Error != nil {
			log.Fatalf("failed to create stock9: %s\n", result.Error)
		}

		result = db.Create(&stock10)

		if result.Error != nil {
			log.Fatalf("failed to create stock10: %s\n", result.Error)
		}

		result = db.Create(&stock11)

		if result.Error != nil {
			log.Fatalf("failed to create stock11: %s\n", result.Error)
		}

		result = db.Create(&stock12)

		if result.Error != nil {
			log.Fatalf("failed to create stock12: %s\n", result.Error)
		}

		log.Println("12 stocks created")
	}
}
