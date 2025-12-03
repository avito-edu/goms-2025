// db.go
package db

import "database/sql"

type User struct {
	ID   int
	Name string
}

func GetUser(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.
		QueryRow("SELECT id, name FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
