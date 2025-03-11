package models

type User struct {
	ID        int
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Nickname  string `db:"nickname"`
}
