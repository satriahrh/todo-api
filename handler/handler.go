package handler

import (
	"encoding/json"
	"net/http"

	"github.com/satriahrh/todo-api/dto"
	"github.com/satriahrh/todo-api/usecase"
)

type handler struct {
	usecaseTodoitem usecase.TodoItemUsecase
}

func NewHandler(usecaseTodoitem usecase.TodoItemUsecase) *handler {
	return &handler{
		usecaseTodoitem: usecaseTodoitem,
	}
}

func (h *handler) Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req dto.CreateItemRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Write([]byte("ga bisa decode json nya"))
			w.WriteHeader(422)
			return
		}

		err = h.usecaseTodoitem.CreateItem(r.Context(), req)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
			return
		}

		w.Write([]byte("Sukses"))
	}
}
