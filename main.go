package main

import (
	"log"
	"net/http"

	"github.com/satriahrh/todo-api/handler"
	"github.com/satriahrh/todo-api/repository/database"
	todoitemusecase "github.com/satriahrh/todo-api/usecase/todoitem"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:rootpw@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Panicf("Cannot open database: %s\n", err.Error())
	}
	repositoryDatabase := database.New(db)

	usecaseTodoitem := todoitemusecase.New(repositoryDatabase)

	handler := handler.NewHandler(usecaseTodoitem)

	http.HandleFunc("/todo", handler.Todo)
	log.Println("Jalan di localhost:8000")

	log.Fatal(
		http.ListenAndServe(":8000", nil),
	)
}
