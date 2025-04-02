package store

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PostgresUserStore struct{
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore{
	return &PostgresUserStore{
		db: db,
	}
}


type password struct{
	plaintext *string
	hash []byte
}

func(p *password) Set(plaintextPassword string) error{
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil{
		return fmt.Errorf("error hashing password: %w", err)
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}


type User struct{
	ID int `json:"id"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	PasswordHash password `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type UserStore interface{
	CreateUser (*User) error
}

func (s *PostgresUserStore) CreateUser(user *User) error {
	query := `
		INSERT INTO users(fullname, email, password)
		VALUES ($1, $2, $3)	
		RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRow(query, user.FullName, user.Email, user.PasswordHash.hash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt )
	if err != nil{
		return err
	}
	return nil
}


