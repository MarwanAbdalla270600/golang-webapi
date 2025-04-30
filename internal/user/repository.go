package user

import (
	"database/sql"
	"fmt"
)

func StoreUserInDatabase(db *sql.DB, user *User) error {
	query := `
		INSERT INTO users (username, password, firstname, lastname)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := db.QueryRow(query, user.Username, user.Password, user.Firstname, user.Lastname).Scan(&user.Id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}
