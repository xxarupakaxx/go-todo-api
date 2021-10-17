package infra

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
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
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}