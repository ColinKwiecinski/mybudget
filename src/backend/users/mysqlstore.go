package users

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStore struct {
	db *sql.DB
}

func NewMysqlStore(DB *sql.DB) *MysqlStore {
	return &MysqlStore{
		db: DB,
	}
}

func (sqlStore *MysqlStore) GetByID(id int64) (*User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE ID=?"
	row, err := sqlStore.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var u User
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) GetByEmail(email string) (*User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE Email=?"
	row, err := sqlStore.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	var u User
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) GetByContactNum(contactNum string) (*User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE Contact_Num=?"
	row, err := sqlStore.db.Query(query, contactNum)
	if err != nil {
		return nil, err
	}
	u := User{}
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) Insert(user *User) (*User, error) {
	query := "INSERT INTO Users (Name, Email, Contact_Num) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := sqlStore.db.Exec(query, user.Name, user.Email, user.Contact_Num)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return sqlStore.GetByID(id)
}

func (sqlStore *MysqlStore) Update(id int64, updates *Updates) (*User, error) {
	query := "UPDATE Users  SET Name = ?, Email = ?, Contact_Num = ? WHERE id = ?"
	_, err := sqlStore.db.Exec(query, updates.Name, updates.Email, updates.Contact_Num, id)
	if err != nil {
		return nil, errors.New("Error Updating User, Check Update Parameters")
	}
	return sqlStore.GetByID(id)
}

func (sqlStore *MysqlStore) Delete(id int64) error {
	query := "DELETE FROM Users WHERE id = ?"
	_, err := sqlStore.db.Exec(query, id)
	if err != nil {
		return errors.New("Error Deleting Existing User from the database")
	}
	return nil
}
