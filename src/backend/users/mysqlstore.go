package users

import (
	"database/sql"
	"errors"

	"mybudget/src/backend/auth"
	"mybudget/src/backend/handlers"

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

func (sqlStore *MysqlStore) GetByID(id int64) (*auth.User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE ID=?"
	row, err := sqlStore.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var u auth.User
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) GetByEmail(email string) (*auth.User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE Email=?"
	row, err := sqlStore.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	var u auth.User
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) GetByContactNum(contactNum string) (*auth.User, error) {
	query := "SELECT ID, Name, Email, Contact_Num FROM Users WHERE Contact_Num=?"
	row, err := sqlStore.db.Query(query, contactNum)
	if err != nil {
		return nil, err
	}
	u := auth.User{}
	if row.Next() {
		row.Scan(&u.ID, &u.Name, &u.Email, &u.Contact_Num)
	}
	return &u, err
}

func (sqlStore *MysqlStore) Insert(user *auth.User) (*auth.User, error) {
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

func (sqlStore *MysqlStore) Update(id int64, updates *auth.Updates) (*auth.User, error) {
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

func (sqlStore *MysqlStore) InsertTransaction(t *handlers.Transaction) error {
	query := "INSERT INTO TRANSACTIONS (User_ID, Transaction_Name, Memo, Transaction_Date, Amount, Transaction_Type) values (?,?,?,?,?,?)"
	_, err := sqlStore.db.Exec(query, t.UID, t.Name, t.Memo, t.Amount, t.Type)
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

func (sqlStore *MysqlStore) GetTransactions(selector string, value string) (*[]auth.Transaction, error) {
	query := "SELECT * FROM Transactions WHERE ? = ?"
	result, err := sqlStore.db.Query(query, selector, value)
	defer result.Close()
	if err != nil {
		return nil, errors.New("error while getting transactions")
	}
	output := make([]handlers.Transaction, 0)
	for result.Next() {
		var temp handlers.Transaction
		if err := result.Scan(&temp.ID, &temp.UID, &temp.Name, &temp.Memo, &temp.Date, &temp.Amount, &temp.Type); err != nil {
			return nil, errors.New("error while getting transactions")
		}
		output = append(output, temp)
	}
	return output, nil
}
