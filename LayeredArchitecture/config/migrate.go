package config

import (
	"github.com/jinzhu/gorm"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	conn.AutoMigrate(domain.News{}, domain.Topic{})
	log.Println("Migration has been processed")

	return conn, err
}
