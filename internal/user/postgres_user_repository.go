package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(host string, port int, user, password, dbname string) (*PostgresUserRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresUserRepository{db: db}, nil
}

// Implement UserRepository interface methods here

func (r *PostgresUserRepository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Implement other UserRepository interface methods
