package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	UseCases UseCase
}

func NewHandler(router *mux.Router, usecase UseCase) {
	handler := &Handler{
		UseCases: usecase,
	}

	router.HandleFunc("/", handler.Home).Methods(http.MethodGet)
	router.HandleFunc("/palindrome", handler.Palindrome).Methods(http.MethodGet)
	router.HandleFunc("/languages/Nomor-3", handler.Languages).Methods(http.MethodGet)
	router.HandleFunc("/languages", handler.GetAllLanguages).Methods(http.MethodGet)
	router.HandleFunc("/language", handler.AddLanguage).Methods(http.MethodPost)
	router.HandleFunc("/language/{id}", handler.GetLanguage).Methods(http.MethodGet)
	router.HandleFunc("/language/{id}", handler.UpdateLanguage).Methods(http.MethodPatch)
	router.HandleFunc("/language/{id}", handler.DeleteLanguage).Methods(http.MethodDelete)
}

func (handler *Handler) Home(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	msg := "Hello Go Developers"
	res = response.Success(response.StatusOK, msg)

	res.JSON(w)
}

func (handler *Handler) Palindrome(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput models.Palindrome

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil)
		res.JSON(w)
		return
	}

	res = handler.UseCases.Palindrome(userInput.Data)

	res.JSON(w)
}

func (handler *Handler) Languages(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	res = handler.UseCases.Language()

	res.JSON(w)
}

func (handler *Handler) AddLanguage(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var language models.ProgrammingLanguage

	if err := json.NewDecoder(r.Body).Decode(&language); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil)
		res.JSON(w)
		return
	}

	res = handler.UseCases.AddLanguage(language)

	res.JSON(w)
}

func (handler *Handler) GetLanguage(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res = response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil)
		res.JSON(w)
		return
	}

	res = handler.UseCases.GetLanguage(id)

	res.JSON(w)
}

func (handler *Handler) GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	res = handler.UseCases.GetAllLanguages()

	res.JSON(w)
}

func (handler *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput models.ProgrammingLanguage

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res = response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil)
		res.JSON(w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		res = response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil)
		res.JSON(w)
		return
	}

	res = handler.UseCases.UpdateLanguage(id, userInput)

	res.JSON(w)
}

func (handler *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res = response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil)
		res.JSON(w)
		return
	}

	res = handler.UseCases.DeleteLanguage(id)

	res.JSON(w)
}
