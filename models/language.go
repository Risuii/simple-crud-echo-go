package models

type ProgrammingLanguage struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"fuctional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}
