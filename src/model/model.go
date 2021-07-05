package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConnect() (db *sql.DB) {
	err:=godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbDriver:="mysql"
	dbOption:="?parseTime=true"
	db,err=sql.Open(dbDriver,os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	return db
}