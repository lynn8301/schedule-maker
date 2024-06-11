package dao

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

const DRIVER = "mysql"

type Config struct {
	Url      string `json: url`
	Username string `json: username`
	Password string `json: passowrd`
	DbName   string `json: dbName`
	Port     string `json: port`
}

func (c *Config) getConfig() *Config {
	jsonFile, err := os.ReadFile("config/local.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(jsonFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func IintMySQL(config *Config) (*grom.DB, err error) {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Url,
		config.Port,
		config.DbName)

	// connect db
	db, err := gorm.Open(DRIVER, databaseURI)
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDB(db *grom.DB) {
	return db.Close()
}
