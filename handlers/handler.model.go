package handlers

import "github.com/Rahul-aithal/ThumbPicker/db"
type  handler struct {
	q *db.Queries
}

func NewHandler(q *db.Queries) *handler {
	return &handler{q: q}
}
