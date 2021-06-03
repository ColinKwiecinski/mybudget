package auth

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStore struct {
	db *sql.DB
}

func NewMysqlStore(dsn string) *MysqlStore {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("error opening database")
		return nil
	}
	log.Println("Opened DB Successfully")
	mysqlStore := &MysqlStore{}
	mysqlStore.db = db
	return mysqlStore
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
	query := "INSERT INTO Users (Name, Email, Contact_Num) VALUES (?, ?, ?)"
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
		return nil, errors.New("error Updating User, Check Update Parameters")
	}
	return sqlStore.GetByID(id)
}

func (sqlStore *MysqlStore) Delete(id int64) error {
	query := "DELETE FROM Users WHERE id = ?"
	_, err := sqlStore.db.Exec(query, id)
	if err != nil {
		return errors.New("error Deleting Existing User from the database")
	}
	return nil
}

func (sqlStore *MysqlStore) InsertTransaction(t *Transaction) error {
	query := "INSERT INTO TRANSACTIONS (User_ID, Transaction_Name, Memo, Transaction_Date, Amount, Transaction_Type) values (?,?,?,?,?,?)"
	_, err := sqlStore.db.Exec(query, t.UID, t.Name, t.Memo, t.Date, t.Amount, t.Type)
	if err != nil {
		return errors.New("error while inserting new transaction")
	}
	return nil
}

func (sqlStore *MysqlStore) DeleteTransaction(id int64) error {
	query := "DELETE FROM TRANSACTIONS WHERE ID = ?"
	_, err := sqlStore.db.Exec(query, id)
	if err != nil {
		return errors.New("error while deleting transaction")
	}
	return nil
}

func (sqlStore *MysqlStore) GetTransactions(selector string, value string, uid int64) (*[]Transaction, error) {
	query := "SELECT * FROM Transactions WHERE ? = ? AND uid = ?"
	result, err := sqlStore.db.Query(query, selector, value, uid)
	if err != nil {
		return nil, errors.New("error while getting transactions")
	}
	output := make([]Transaction, 0)
	for result.Next() {
		var temp Transaction
		if err := result.Scan(&temp.ID, &temp.UID, &temp.Name, &temp.Memo, &temp.Date, &temp.Amount, &temp.Type); err != nil {
			return nil, errors.New("error while getting transactions")
		}
		output = append(output, temp)
	}
	defer result.Close()
	return &output, nil
}
