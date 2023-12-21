package models

type Collection struct {
	ID       int64  `json:"id"`
	LaidDate string `json:"laidDate"`
	Number   int    `json:"number"`
}
