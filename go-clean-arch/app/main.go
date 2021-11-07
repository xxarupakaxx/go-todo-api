package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	_articleHttpDelivery "github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/delivery/http"
	_articleHttpDeliveryMiddleware "github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/delivery/http/middleware"
	_articleRepo "github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/repository/mysql"
	_articleUcase "github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/usecase"
	_authorRepo "github.com/xxarupakaxx/go-todo-api/go-clean-arch/author/repository/mysql"
	"log"
	"net/url"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	if err = dbConn.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Success")
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}()
	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	a := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt64("context.timeout")) * time.Second
	au := _articleUcase.NewArticleUsecase(a, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
