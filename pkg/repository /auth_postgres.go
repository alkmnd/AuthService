package repository

import (
	"authService/pkg/repository /models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(email string) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf("INSERT INTO %s (email) values ($1) RETURNING id", "users")
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateToken(token models.Token) error {
	query := fmt.Sprintf("INSERT INTO %s (token_hash, user_id, ip_address, jti, created_at, expired_at, is_revoked) values ($1, $2, $3, $4, $5, $6, $7)", "tokens")
	_, err := r.db.Exec(query, token.TokenHash, token.UserId, token.IpAddress, token.Jti, token.CreatedAt, token.ExpiresAt, token.IsRevoked)
	return err
}
