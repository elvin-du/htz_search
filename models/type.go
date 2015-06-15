package models

import (
	"time"
)

type Article struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Abs     string    `json:"abs"`
	Kind    string    `json:"kind"`
	Author  string    `json:"kind"`
	Ctime   time.Time `json:"ctime"`
}
