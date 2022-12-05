package repository

import (
	"database/sql"

	"github.com/warete/edu-crud/pkg/common/db"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	SecondName string `json:"secondName"`
}

type UserRepository struct {
	conn *db.Connection
}

func InitUserRepository(c *db.Connection) *UserRepository {
	return &UserRepository{
		c,
	}
}

func (r *UserRepository) GetAll() ([]User, error) {
	rows, err := r.conn.DB.Query("select id, name, last_name, second_name from users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		tmpUser := User{}
		err = rows.Scan(&tmpUser.Id, &tmpUser.Name, &tmpUser.LastName, &tmpUser.SecondName)
		if err != nil {
			return nil, err
		}

		users = append(users, tmpUser)
	}
	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetById(id int) (User, error) {
	stmt, err := r.conn.DB.Prepare("select id, name, last_name, second_name from users where id = ?")

	if err != nil {
		return User{}, nil
	}

	user := User{}

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.LastName, &user.SecondName)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) Add(newUser User) (User, error) {
	tx, err := r.conn.DB.Begin()
	if err != nil {
		return User{}, err
	}

	stmt, err := tx.Prepare("insert into users (name, last_name, second_name) values (?, ?, ?)")
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(newUser.Name, newUser.LastName, newUser.SecondName)

	if err != nil {
		return User{}, err
	}

	tx.Commit()

	lastId, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	newUser.Id = int(lastId)

	return newUser, nil
}

func (r *UserRepository) Update(id int, user User) error {
	tx, err := r.conn.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("update users set name = ?, last_name = ?, second_name =? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.LastName, user.SecondName, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *UserRepository) Delete(id int) error {
	tx, err := r.conn.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
