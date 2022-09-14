package models

type Hacker struct {
	Name  string `json:"name,omitempty"`
	Score int    `json:"score,omitempty"`
}
