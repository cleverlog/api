package tools

import "github.com/google/uuid"

func StringToUUID(str string) uuid.UUID {
	id, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil
	}

	return id
}