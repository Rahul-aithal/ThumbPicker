package common

import (
	"log"

	"github.com/google/uuid"
)

func GetUUID(id string) uuid.UUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Println(err, "Error while converting to uuid")
		panic("Error while converting to uuid")
	}
	return uid
}
