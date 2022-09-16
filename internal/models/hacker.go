package models

type Hacker struct {
	Name  string  `json:"name,omitempty"`
	Score float64 `json:"score,omitempty"`
}
