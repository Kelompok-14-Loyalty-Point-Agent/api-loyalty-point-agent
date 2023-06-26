package mysql_driver

import (
	"api-loyalty-point-agent/drivers/mysql/profiles"
	"api-loyalty-point-agent/drivers/mysql/providers"
	"api-loyalty-point-agent/drivers/mysql/redeems"
	"api-loyalty-point-agent/drivers/mysql/stock_details"
	"api-loyalty-point-agent/drivers/mysql/stock_transactions"
	"api-loyalty-point-agent/drivers/mysql/stocks"
	"api-loyalty-point-agent/drivers/mysql/transactions"
	"api-loyalty-point-agent/drivers/mysql/users"
	"api-loyalty-point-agent/drivers/mysql/vouchers"

	"time"

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

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&profiles.Profile{}, &users.User{}, &providers.Provider{}, &stocks.Stock{}, &stock_details.StockDetail{}, &stock_transactions.StockTransaction{}, &transactions.Transaction{}, &vouchers.Voucher{}, &redeems.Redeem{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func SeedAdmin(db *gorm.DB) {
	var admin = users.User{
		Name:      "admin",
		Password:  "admin123",
		Email:     "admin@example.com",
		Role:      "admin",
		ProfileID: 0,
		Profile:   profiles.Profile{},
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
		profile.Name = admin.Name

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

func SeedCustomer(db *gorm.DB) {
	var customer = users.User{
		Name:      "customer_example",
		Password:  "customer123",
		Email:     "customer_example@example.com",
		Role:      "customer",
		ProfileID: 0,
		Profile:   profiles.Profile{},
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("failed to hash customer password: %s\n", err)
	}

	customer.Password = string(password)

	var record users.User

	_ = db.First(&record, "email = ?", customer.Email)

	if record.ID != 0 {
		log.Printf("customer already exists\n")
	} else {
		filePath := "./assets/users/user.png"

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to create customer: %s\n", err.Error())
		}

		defer file.Close()

		url, err := aws_driver.UploadFileToBucket("user.png", file)
		if err != nil {
			log.Fatalf("failed to create customer: %s\n", err.Error())
		}

		err = aws_driver.DownloadFileFromBucket(url, "./assets/users/")
		if err != nil {
			log.Fatalf("failed to create customer: %s\n", err.Error())
		}

		var profile profiles.Profile

		profile.URL = url
		profile.Gender = "man"
		profile.Address = "Summarecon BSD"
		profile.Age = 25
		profile.Member = "gold"
		profile.MonthlyTransaction = 10
		profile.TransactionMade = 10
		profile.Point = float64(1000000)
		profile.Phone = "082189998231"
		profile.Name = customer.Name

		result := db.Create(&profile)

		if result.Error != nil {
			log.Fatalf("failed to create customer's profile: %s\n", result.Error)
		}

		customer.ProfileID = profile.ID
		customer.Profile = profile

		result = db.Create(&customer)

		if result.Error != nil {
			log.Fatalf("failed to create customer: %s\n", result.Error)
		}

		log.Println("customer created")
	}
}

func SeedProvider(db *gorm.DB) {
	providersData := []providers.Provider{
		{Name: "Telkomsel", URL: ""},
		{Name: "XL", URL: ""},
		{Name: "Smartfren", URL: ""},
		{Name: "Indosat Ooredoo", URL: ""},
		{Name: "Axis", URL: ""},
		{Name: "3", URL: ""},
	}

	providerImages := []string{
		"./assets/providers/tsel.png",
		"./assets/providers/xl.png",
		"./assets/providers/smartfren.png",
		"./assets/providers/indosat.png",
		"./assets/providers/axis.png",
		"./assets/providers/3.png",
	}

	var record providers.Provider
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("provider already exists\n")
	} else {
		for i, provider := range providersData {
			filePath := providerImages[i]

			file, err := os.Open(filePath)
			if err != nil {
				log.Fatalf("failed to create provider: %s\n", err.Error())
			}
			defer file.Close()

			url, err := aws_driver.UploadFileToBucket(fmt.Sprintf("%s.png", provider.Name), file)
			if err != nil {
				log.Fatalf("failed to create provider: %s\n", err.Error())
			}

			provider.URL = url

			result := db.Create(&provider)

			if result.Error != nil {
				log.Fatalf("failed to create provider: %s\n", err.Error())
				break
			}
		}

		log.Printf("%d providers created\n", len(providersData))
	}
}

func SeedStockDetail(db *gorm.DB) {
	// Data stock details
	stockDetailData := []stock_details.StockDetail{
		{StockID: 2, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 2, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 2, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 2, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 2, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 1, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 1, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 1, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 1, Stock: 25, Price: 102500, Quantity: 100},

		{StockID: 4, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 4, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 4, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 4, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 4, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 3, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 3, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 3, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 3, Stock: 25, Price: 102500, Quantity: 100},

		{StockID: 6, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 6, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 6, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 6, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 6, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 5, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 5, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 5, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 5, Stock: 25, Price: 102500, Quantity: 100},
		
		{StockID: 8, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 8, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 8, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 8, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 8, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 7, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 7, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 7, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 7, Stock: 25, Price: 102500, Quantity: 100},

		{StockID: 10, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 10, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 10, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 10, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 10, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 9, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 9, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 9, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 9, Stock: 25, Price: 102500, Quantity: 100},
		
		{StockID: 12, Stock: 5000, Price: 9500, Quantity: 100},
		{StockID: 12, Stock: 10000, Price: 14500, Quantity: 100},
		{StockID: 12, Stock: 25000, Price: 27500, Quantity: 100},
		{StockID: 12, Stock: 50000, Price: 52500, Quantity: 100},
		{StockID: 12, Stock: 100000, Price: 102500, Quantity: 100},
		{StockID: 11, Stock: 5, Price: 29500, Quantity: 100},
		{StockID: 11, Stock: 10, Price: 54500, Quantity: 100},
		{StockID: 11, Stock: 15, Price: 77500, Quantity: 100},
		{StockID: 11, Stock: 25, Price: 102500, Quantity: 100},
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// Panggil fungsi SeedStock di dalam transaksi
		if err := SeedStock(tx); err != nil {
			return err
		}

		var record stock_details.StockDetail
		_ = tx.First(&record)

		if record.ID != 0 {
			log.Printf("stock detail already exists\n")
		} else {
			for _, stock := range stockDetailData {
				result := tx.Create(&stock)
				if result.Error != nil {
					return result.Error
				}
			}
			log.Printf("%d stock detail created\n", len(stockDetailData))
		}

		return nil
	})

	if err != nil {
		log.Fatalf("failed to seed stock detail: %s\n", err.Error())
	}
}

func SeedStock(db *gorm.DB) error {
	// Data stocks
	stocksData := []stocks.Stock{
		{Type: "data", TotalStock: 100, ProviderID: 1, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 1, LastTopUp: time.Now()},
		{Type: "data", TotalStock: 100, ProviderID: 2, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 2, LastTopUp: time.Now()},
		{Type: "data", TotalStock: 100, ProviderID: 3, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 3, LastTopUp: time.Now()},
		{Type: "data", TotalStock: 100, ProviderID: 4, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 4, LastTopUp: time.Now()},
		{Type: "data", TotalStock: 100, ProviderID: 5, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 5, LastTopUp: time.Now()},
		{Type: "data", TotalStock: 100, ProviderID: 6, LastTopUp: time.Now()},
		{Type: "credit", TotalStock: 1000000, ProviderID: 6, LastTopUp: time.Now()},
	}

	var record stocks.Stock
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("stock already exists\n")
	} else {
		for _, stock := range stocksData {
			result := db.Create(&stock)
			if result.Error != nil {
				return result.Error
			}
		}

		log.Printf("%d stocks created\n", len(stocksData))
	}

	return nil
}

func SeedVoucher(db *gorm.DB) error {
	// Data vouchers
	vouchersData := []vouchers.Voucher{
		{Product: "Phone Balance", Benefit: "Rp 10.000", Cost: 300},
		{Product: "Internet Data", Benefit: "1GB", Cost: 600},
	}

	var record vouchers.Voucher
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("voucher already exists\n")
	} else {
		for _, voucher := range vouchersData {
			result := db.Create(&voucher)
			if result.Error != nil {
				return result.Error
			}
		}

		log.Printf("%d voucher created\n", len(vouchersData))
	}

	return nil
}

func SeedTransaction(db *gorm.DB) error {
	// Data transactions
	transactionData := []transactions.Transaction{
		{Phone: "08198293845", StockDetailsID: 1, StockID: 1, Price: 9500, Payment_method: "Go Pay", Point: 9.4, Status: "success", Description: "top up", UserID: 2, Product: "Credit Telkomsel"},
		{Phone: "08198293845", StockDetailsID: 2, StockID: 1, Price: 14500, Payment_method: "Dana", Point: 14.4, Status: "success", Description: "top up", UserID: 2, Product: "Credit Telkomsel"},
	}

	var record transactions.Transaction
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("voucher already exists\n")
	} else {
		for _, transaction := range transactionData {
			result := db.Create(&transaction)
			if result.Error != nil {
				return result.Error
			}
		}

		log.Printf("%d voucher created\n", len(transactionData))
	}

	return nil
}
