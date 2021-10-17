package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	err:=godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbDriver:="mysql"
	dbOption:="?parseTime=true"
	conn,err:=sql.Open(dbDriver,os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Ping();err == nil{
		fmt.Println("success")
	}else {
		fmt.Println("fall")
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}