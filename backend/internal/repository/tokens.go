package repository

import (
	"context"
	"github.com/kendoow/SportApp/backend/db"
)

func SaveToken(ctx context.Context, id int64, rToken string) error {
	query := `INSERT INTO tokens(id, token) 
		VALUES ($1, $2)`

	if err := db.GetDB().QueryRow(ctx, query, id, rToken).Scan(); err != nil {
		return err
	}

	return nil
}

func DeleteToken(ctx context.Context, rToken string) error {
	query := `DELETE FROM tokens
			WHERE token = $1;`

	if err := db.GetDB().QueryRow(ctx, query, rToken).Scan(); err != nil {
		return err
	}

	return nil
}

func GetTokenByUserId(ctx context.Context, id int64) (string, error) {
	query := `SELECT token 
			  FROM tokens
			  WHERE id = $1;`

	var token string

	if err := db.GetDB().QueryRow(ctx, query, id).Scan(&token); err != nil {
		return "", err
	}

	return token, nil
}

func UpdateTokenByUserId(ctx context.Context, id int64, refreshToken string) error {
	query := `UPDATE tokens 
			  SET token = $1
			  WHERE id = $2`

	if err := db.GetDB().QueryRow(ctx, query,
		refreshToken,
		id).Scan(); err != nil {
		return err
	}

	return nil
}
