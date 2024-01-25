package models

type Tag struct {
	ID     int
	Name   string
	PostID int
	Count  int
}
