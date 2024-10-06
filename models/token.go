package models

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
	TokenHash string    `json:"token_hash" db:"token_hash"`
	IpAddress string    `json:"ip_address" db:"ip_address"`
	Jti       string    `json:"jti" db:"jti"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}
