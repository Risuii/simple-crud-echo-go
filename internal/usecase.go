package internal

import (
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models"
)

type UseCase struct{}

func NewUseCase() *UseCase {
	return &UseCase{}
}

var languages []models.ProgrammingLanguage

func (uc *UseCase) Palindrome(params string) response.Response {
	var msg string
	for i, j := 0, len(params)-1; i < j; i, j = i+1, j-1 {
		if params[i] != params[j] {
			msg = "Not palindrome"
			return response.Error(response.StatusBadRequest, exception.ErrNotPalindrome, msg)
		}
	}

	msg = "Palindrome"

	return response.Success(response.StatusOK, msg)
}

func (uc *UseCase) Language() response.Response {
	lang := models.ProgrammingLanguage{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: models.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}

	return response.Success(response.StatusOK, lang)
}

func (uc *UseCase) AddLanguage(params models.ProgrammingLanguage) response.Response {
	languages = append(languages, params)

	return response.Success(response.StatusOK, languages)
}

func (uc *UseCase) GetLanguage(id int) response.Response {
	if id >= len(languages) {
		return response.Error(response.StatusNotFound, exception.ErrNotFound, nil)
	}

	return response.Success(response.StatusOK, languages[id])
}

func (uc *UseCase) GetAllLanguages() response.Response {
	return response.Success(response.StatusOK, languages)
}

func (uc *UseCase) UpdateLanguage(id int, params models.ProgrammingLanguage) response.Response {
	if id >= len(languages) {
		return response.Error(response.StatusNotFound, exception.ErrNotFound, nil)
	}

	languages[id] = params

	return response.Success(response.StatusOK, params)
}

func (uc *UseCase) DeleteLanguage(id int) response.Response {
	if id >= len(languages) {
		return response.Error(response.StatusNotFound, exception.ErrNotFound, nil)
	}

	languages = append(languages[:id], languages[id+1:]...)
	return response.Success(response.StatusOK, languages)
}
