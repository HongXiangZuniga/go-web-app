package mysql

import (
	"database/sql"

	"github.com/HongXiangZuniga/login-go/pkg/auth"
)

type storage struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) auth.Repository {
	return &storage{db: db}
}

func (stg *storage) AuthUser(email, password string) (*bool, error) {
	result := false
	count := 0
	query := `
		SELECT 
			COUNT(u.id) 
		FROM 
			User u
		WHERE 
			u.email = ? 
		AND 
			u.password  = ?
	`
	rows, err := stg.db.Query(query, email, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}
	if count == 1 {
		result = true
	}
	return &result, nil
}

func (stg *storage) GetSalt(email string) (*string, error) {
	var salt string
	query := `
			SELECT 
				u.salt 
			FROM 
				User u 
			WHERE 
				u.email = ?
	`
	rows, err := stg.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&salt)
		if err != nil {
			return nil, err
		}
	}
	return &salt, nil
}
