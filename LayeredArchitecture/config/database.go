package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type ConfigDB struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Dbname   string `toml:"dbname"`
}

func (c *ConfigDB) Read() {
	f,err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := toml.NewDecoder(f).Decode(&c); err != nil {
		log.Fatalln(err)
	}
}

var config = ConfigDB{}

func ConnectDB() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}
