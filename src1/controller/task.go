package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/xxarupakaxx/go-todo-api/src1/model1"
	"log"
	"net/http"
	"strconv"
	"time"
)

func TaskGET(c echo.Context) error {
	db:=model.DBConnect()
	result,err:=db.Query("SELECT * FROM task ORDER BY id DESC ")
	if err != nil {
		log.Fatal(err)
	}
	
	tasks:=make([]model.Task,0)
	for result.Next() {
		
		var id uint
		var createdAt,updatedAt time.Time
		var title string
		
		err=result.Scan(&id,&createdAt,&updatedAt,&title)
		if err != nil {
			log.Fatal(err)
		}
		task:=model.Task{
			id,
			createdAt,
			updatedAt,
			title,
		}
		tasks=append(tasks,task)
	}
	return c.JSON(http.StatusOK,tasks)
}

func FindByID(id uint) model.Task {
	db:=model.DBConnect()
	result,err:=db.Query("SELECT * FROM task WHERE id=?",id)
	if err != nil {
		log.Fatal(err)
	}
	task:=model.Task{}
	for result.Next(){
		var createdAt,updatedAt time.Time
		var title string
		err=result.Scan(&id,&createdAt,&updatedAt,&title)
		if err != nil {
			log.Fatal(err)
		}

		task=model.Task{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Title:     title,
		}

	}
	return task

}

func TaskPOST(c echo.Context) error {
	db:=model.DBConnect()

	title:=c.FormValue("title")
	now:=time.Now()

	_,err:=db.Exec("INSERT INTO task (title,created_at,updated_at) VALUES (?,?,?)",title,now,now)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK,fmt.Sprintf("post sent. title:%s",title))
}
func TaskPATCH(c echo.Context) error {
	db:=model.DBConnect()

	id,_:=strconv.Atoi(c.Param("id"))
	title:=c.FormValue("title")
	now:=time.Now()

	_,err:=db.Exec("UPDATE task SET title=?, updated_at=? WHERE id=?",title,now,id)
	if err != nil {
		log.Fatal(err)
	}

	task:=FindByID(uint(id))

	fmt.Println(task)
	return c.JSON(http.StatusOK,task)
}

func TaskDELETE(c echo.Context) error {
	db:=model.DBConnect()

	id,_:=strconv.Atoi(c.Param("id"))
	_,err:=db.Query("DELETE FROM task WHERE id=?",id)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK,"deleted")
}
