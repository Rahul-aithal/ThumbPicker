package types

import "github.com/google/uuid"

type ThumbData struct {
	Path      string
	Index     int
	Id        uuid.UUID
	TimeStamp string
}
