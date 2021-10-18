package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"log"
)

type ConfigDB struct {
	User string
	Password string
	Host string
	Port string
	Dbname string
}

func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatalln(err)
	}
}

var config  = ConfigDB{}

func ConnectDB() (*gorm.DB,error) {
	config.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)

	db,err := gorm.Open("mysql",dsn)
	if err != nil {
		return nil,err
	}
	return db,err
}