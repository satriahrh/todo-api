package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (h *handler) Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req dto.CreateItemRequest

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
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
	default:
		json.NewEncoder(w).Encode(
			Response{Message: "API Tidak valid"},
		)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

func (h *handler) TodoWithFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	name := r.FormValue("name")
	description := r.FormValue("description")
	fmt.Println(name, description)
	file, header, err := r.FormFile("document")
	if err != nil {
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.usecaseTodoitem.CreateItemWithFile(r.Context(), dto.CreateItemWithFileRequest{
		CreateItemRequest: dto.CreateItemRequest{
			Name:        name,
			Description: description,
		},
		Document: fileBytes,
	})
	if err != nil {
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Message: header.Filename,
	})
	w.WriteHeader(http.StatusOK)
}
