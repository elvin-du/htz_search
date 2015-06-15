package controllers

type Article struct {
	AID     int    `json:"aid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
