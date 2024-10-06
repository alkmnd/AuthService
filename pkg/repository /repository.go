package repository

import (
	"authService/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(email string) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf("INSERT INTO %s (email) values ($1) RETURNING id", "users")
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *Repository) GetUser(userId uuid.UUID) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, email FROM %s WHERE id=$1", "users")
	err := r.db.Get(&user, query, userId)
	return user, err
}

func (r *Repository) CreateToken(token models.Token) error {
	query := fmt.Sprintf("INSERT INTO %s (token_hash, user_id, ip_address, jti, created_at, expired_at) values ($1, $2, $3, $4, $5, $6)", "tokens")
	_, err := r.db.Exec(query, token.TokenHash, token.UserId, token.IpAddress, token.Jti, token.CreatedAt, token.ExpiresAt)
	return err
}

func (r *Repository) GetTokenInfo(tokenHash string, userId uuid.UUID) (models.Token, error) {
	var token models.Token
	query := fmt.Sprintf("SELECT user_id, token_hash, ip_address, jti, created_at, expires_at FROM %s WHERE token_hash=$1 AND user_id=$2", "tokens")
	err := r.db.Get(&token, query, tokenHash, userId)
	return token, err
}

func (r *Repository) UpdateAccessId(tokenHash string, userId uuid.UUID, jti string) error {
	query := fmt.Sprintf("UPDATE %s SET jti = $1 WHERE token_hash = $2 AND user_id=$3", "tokens")

	_, err := r.db.Exec(query, jti, tokenHash, userId)

	return err
}
