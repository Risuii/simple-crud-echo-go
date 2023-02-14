package models

type Relation struct {
	InfluencedBy []string `json:"influence-by"`
	Influences   []string `json:"influences"`
}
