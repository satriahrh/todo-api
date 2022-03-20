package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/satriahrh/todo-api/handler"
	"github.com/satriahrh/todo-api/repository/database"
	"github.com/satriahrh/todo-api/repository/gcs"
	todoitemusecase "github.com/satriahrh/todo-api/usecase/todoitem"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")))
	if err != nil {
		log.Panicf("Cannot open database: %s\n", err.Error())
	}
	repositoryDatabase := database.New(db)
	repositoryGcs := gcs.New()

	usecaseTodoitem := todoitemusecase.New(repositoryDatabase, repositoryGcs)

	handler := handler.NewHandler(usecaseTodoitem)

	http.HandleFunc("/todo", handler.Todo)
	http.HandleFunc("/todo-with-file", handler.TodoWithFile)
	log.Println("Jalan di localhost:8000")

	log.Fatal(
		http.ListenAndServe(":8000", nil),
	)
}
