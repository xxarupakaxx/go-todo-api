package config

import (
	"github.com/jinzhu/gorm"
	"log"
)

func DBMigrate() (*gorm.DB, error) {
	conn,err := ConnectDB()
	if err != nil {
		return nil, err
		}
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	conn.AutoMigrate(domain.News{},domain.Topic{})
	log.Println("Migration has been processed")

	return conn,err
}