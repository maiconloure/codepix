package db

import (
	"github.com/maiconloure/codepix/domain/model"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	-"github.com/lib/pq"
	-"gorm.io/driver/sqlite"
)

func init() {
	_, b, _, := runtime.Caller(skip:0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf(format: "Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv(key: "dsn")
		db, err = gorm.Open(os.Getenv(key: "dbType"), dsn)
	} else {
		dsn = os.Getenv(key: "dsnTest")
		db, err = gorm.Open(os.Getenv(key: "dbTypeTest"), dsn)
	}

	if err != nil {
		log.Fatalf(format: "Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv(key: "debug") == "true" {
		db.LogMode(enable: true)
	}

	if os.Getenv(key: "AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}