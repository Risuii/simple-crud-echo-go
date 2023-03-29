package internal

import (
	"strconv"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCases UseCase
}

func NewHandler(e *echo.Echo, usecase UseCase) {
	handler := &Handler{
		UseCases: usecase,
	}

	e.GET("/", handler.Home)
	e.GET("/palindrome", handler.Palindrome)
	e.GET("/languages/Nomor-3", handler.Languages)
	e.GET("/languages", handler.GetAllLanguages)
	e.POST("/language", handler.AddLanguage)
	e.GET("/language/:id", handler.GetLanguage)
	e.PATCH("/language/:id", handler.UpdateLanguage)
	e.DELETE("/language/:id", handler.DeleteLanguage)
}

func (handler *Handler) Home(c echo.Context) error {
	msg := "Hello Go Developers"
	return response.Success(response.StatusOK, msg).JSON(c.Response())
}

func (handler *Handler) Palindrome(c echo.Context) error {
	var userInput models.Palindrome

	if err := c.Bind(&userInput); err != nil {
		return response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil).JSON(c.Response())
	}

	return handler.UseCases.Palindrome(userInput.Data).JSON(c.Response())
}

func (handler *Handler) Languages(c echo.Context) error {
	return handler.UseCases.Language().JSON(c.Response())
}

func (handler *Handler) AddLanguage(c echo.Context) error {
	var language models.ProgrammingLanguage

	if err := c.Bind(&language); err != nil {
		return response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil).JSON(c.Response())
	}

	return handler.UseCases.AddLanguage(language).JSON(c.Response())
}

func (handler *Handler) GetLanguage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil).JSON(c.Response())
	}

	return handler.UseCases.GetLanguage(id).JSON(c.Response())
}

func (handler *Handler) GetAllLanguages(c echo.Context) error {
	return handler.UseCases.GetAllLanguages().JSON(c.Response())
}

func (handler *Handler) UpdateLanguage(c echo.Context) error {
	var userInput models.ProgrammingLanguage

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil).JSON(c.Response())
	}

	if err := c.Bind(&userInput); err != nil {
		return response.Error(response.StatusUnprocessableEntity, exception.ErrUnprocessableEntity, nil).JSON(c.Response())
	}

	return handler.UseCases.UpdateLanguage(id, userInput).JSON(c.Response())
}

func (handler *Handler) DeleteLanguage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer, nil).JSON(c.Response())
	}

	return handler.UseCases.DeleteLanguage(id).JSON(c.Response())
}
