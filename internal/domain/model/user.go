package model

import "github.com/google/uuid"

type User struct {
	UUID                 uuid.UUID
	TelegramLogin        string
	FatsecretAccessToken *string
	FatsecretTokenSecret *string
}
