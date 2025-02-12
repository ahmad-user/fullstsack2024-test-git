package config

import (
	"Encode/model/user"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var JWTSECRET string

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

func assignEnv(c *AppConfig) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(" Gagal memuat file .env, pastikan file .env ada di root proyek!")
	}
	c.DBUsername = os.Getenv("DBUsername")
	c.DBPassword = os.Getenv("DBPassword")
	c.DBPort = os.Getenv("DBPort")
	c.DBHost = os.Getenv("DBHost")
	c.DBName = os.Getenv("DBName")
	JWTSECRET = os.Getenv("JWT_SECRET")
}

func InitConfig() AppConfig {
	var result AppConfig
	assignEnv(&result)
	return result
}

func InitSQL(c AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		c.DBHost, c.DBUsername, c.DBPassword, c.DBName, c.DBPort)

	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ Terjadi error:", err.Error())
		return nil
	}

	db.AutoMigrate(&user.MyClient{})
	fmt.Println("Database berhasil terhubung!")

	return db
}
