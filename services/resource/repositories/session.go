package repositories

import (
	"github.com/exbotanical/gouache/models"
)

func (db *DB) CreateUser(t *models.NewUserTemplate) (string, error) {
	sql := `
    	INSERT INTO user_record
			(
				username,
				passhash
			)
			VALUES ($1, $2)
			RETURNING username;
  	`

	var username string
	row := db.QueryRow(sql,
		t.Username,
		t.Password,
	)
	if err := row.Scan(&username); err != nil {
		return "", err
	}

	return username, nil
}

func (db *DB) GetUser(username string) (models.User, error) {
	sql := `
		SELECT
			username,
			passhash
		FROM user_record
		WHERE username=$1
		LIMIT 1;
	`

	row := db.QueryRow(sql, username)

	var u models.User

	if err := row.Scan(
		&u.Username,
		&u.Password,
	); err != nil {
		return models.User{}, err
	}

	return u, nil
}
