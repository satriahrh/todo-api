package main

import (
	"log"
	"net/http"

	"github.com/satriahrh/todo-api/handler"
	todoitemusecase "github.com/satriahrh/todo-api/usecase/todoitem"
)

func main() {
	usecaseTodoitem := todoitemusecase.New()

	handler := handler.NewHandler(usecaseTodoitem)

	http.HandleFunc("/todo", handler.Todo)

	log.Fatal(
		http.ListenAndServe(":8000", nil),
	)
}
