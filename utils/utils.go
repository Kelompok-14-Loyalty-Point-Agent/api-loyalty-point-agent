package utils

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"math"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}

func GenerateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	fileName := strings.TrimSuffix(originalName, ext)

	timestamp := time.Now().Unix()
	fileName = fileName + "_" + strconv.FormatInt(timestamp, 10)

	uniqueID := uuid.New().String()
	fileName = fileName + "_" + uniqueID

	fileName = fileName + ext

	return fileName
}

func RoundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}


