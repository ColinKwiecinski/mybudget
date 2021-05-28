package auth

// import (
// 	"database/sql"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"strings"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// //ErrUserNotFound is returned when the user can't be found
// var ErrUserNotFound2 = errors.New("user not found")

// //Store represents a store for Users
// type MySQLStore struct {
// 	Database *sql.DB
// }

// func NewSQLStore(dsn string) *MySQLStore {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("error opening database")
// 		return nil
// 	}
// 	log.Println("Opened DB Successfully")
// 	mysqlStore := &MySQLStore{}
// 	mysqlStore.Database = db
// 	return mysqlStore
// }

// //Implementation of GetByID from store.go using MySQL database
// func (db *MySQLStore) GetByID(id int64) (*User, error) {
// 	user := &User{}
// 	err := db.Database.QueryRow("select id, email, passhash, username, name from Users where id = ?", id).Scan(&user.ID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName, &user.PhotoURL)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return user, nil
// 	}
// }

// //Similar to GetByID but check for the email parameter rather than ID
// func (db *MySQLStore) GetByEmail(email string) (*User, error) {
// 	user := &User{}
// 	err := db.Database.QueryRow("select id, email, passhash, username, firstname, lastname, photourl from Users where email = ?", email).Scan(&user.ID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName, &user.PhotoURL)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return user, nil
// 	}
// }

// //Again, same method as GetByEmail and GetByID, using username for the parameter instead to get the user
// func (db *MySQLStore) GetByUserName(username string) (*User, error) {
// 	user := &User{}
// 	err := db.Database.QueryRow("select id, email, passhash, username, firstname, lastname, photourl from Users where username = ?", username).Scan(&user.ID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName, &user.PhotoURL)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return user, nil
// 	}
// }

// //For Insert, create a insert statement with '?' so that it guards against SQL injections.
// //Then using that to insert a new row into the database.
// //Then get the ID of the last inserted row and assign it to the User and return it if there isn't
// //any errors
// func (db *MySQLStore) Insert(user *User) (*User, error) {
// 	insq := "insert into Users(email, passhash, username, firstname, lastname, photourl) values (?, ?, ?, ?, ?, ?)"
// 	res, err := db.Database.Exec(insq, user.Email, user.PassHash, user.UserName, user.FirstName, user.LastName, user.PhotoURL)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		id, err := res.LastInsertId()
// 		if err != nil {
// 			return nil, err
// 		} else {
// 			user.ID = id
// 			return user, nil
// 		}
// 	}
// }

// func (db *MySQLStore) InsertLogIn(user *User, r *http.Request) (*LogIn, error) {
// 	insq := "insert into Logs(userid, logindatetime, ipaddress) values (?, ?, ?)"
// 	ipaddress := ""
// 	if r.Header.Get("X-Forwarded-For") != "" {
// 		ipaddress = strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
// 	} else {
// 		ipaddress = r.RemoteAddr
// 	}
// 	_, err := db.Database.Exec(insq, user.ID, time.Now(), ipaddress)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		login := &LogIn{}
// 		login.IPAdress = ipaddress
// 		login.LoginDateTime = time.Now()
// 		login.UserID = user.ID
// 		return login, nil
// 	}
// }

// //Apply the updates to the MySQL database first then after some error handling,
// //retreive the user by a query to save into the User struct. Then returning the updated user.
// func (db *MySQLStore) Update(id int64, update *Updates) (*User, error) {
// 	_, err := db.Database.Exec("update Users set firstname = ?, lastname = ? where id = ?", update.FirstName, update.LastName, id)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		user := &User{}
// 		err := db.Database.QueryRow("select id, email, passhash, username, firstname, lastname, photourl from Users where id = ?", id).Scan(&user.ID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName, &user.PhotoURL)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return user, nil
// 	}
// }

// //Similar to Update, delete the row from the database first and if an error is returned, then return the error
// //else return nil to show that it has deleted successfully.
// func (db *MySQLStore) Delete(id int64) error {
// 	_, err := db.Database.Exec("delete from Users where id = ?", id)
// 	if err != nil {
// 		return err
// 	} else {
// 		return nil
// 	}
// }
