package sqlrepo

import (
	"fmt"
	"log"
	"database/sql"

	// "github.com/tjeerddie/basic-go-api/repository"
	"github.com/tjeerddie/basic-go-api/entities"
)

var (
	listQuery   = `SELECT * FROM user`
	singleQuery = `SELECT * FROM user WHERE id=?`
	createQuery = `INSERT INTO user (email, password) VALUES (?,?)`
	updateQuery = `UPDATE user SET (?) WHERE id=?`
	deleteQuery = `DELETE FROM user WHERE id=?`
)

type Repository struct {
	db     *sql.DB
}

// Users returns all the users
func (r *Repository) Users() ([]entities.User, error) {
	rows, err := r.db.Query(listQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id       int
		email    string
		password string
		userList []entities.User
	)

	for rows.Next() {
		err = rows.Scan(&id, &email, &password)
		userList = append(userList, entities.User{
			Id:       id,
			Email:    email,
			Password: password,
		})
	}

	return userList, err
}

// UserSingle returns one user by id
func (r *Repository) UserSingle(id int) (*entities.User, error) {
	var user entities.User

	row := r.db.QueryRow(singleQuery, id)
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found: %s", id)
		}

		log.Fatal("Database connection failed: %s", err.Error())

	}
	return &user, nil
}

func (r *Repository) UserStore(user *entities.User) (error) {
	stmt, err := r.db.Prepare(createQuery)
	if err != nil {
		panic(err.Error())
	}

    if res, err := execAffectingOneRow(stmt, &user.Email, &user.Password); err != nil {
        return fmt.Errorf(err.Error())
    } else {
        id, err := res.LastInsertId()
        if err != nil {
            return fmt.Errorf(err.Error())
		}
		user.Id = int(id)
	}

	return nil
}


// New returns a new repository
func New(databaseURL string) *Repository {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	return &Repository{db: db}
}

// Close closes all that needs to close in the repository
func (r *Repository) Close() {
	r.db.Close()
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func execAffectingOneRow(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.Exec(args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}
