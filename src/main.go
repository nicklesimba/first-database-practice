//A third branch just to make sure i really got it down :]

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "bobby2643"
	dbname   = "calhounio_demo"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Returning Data (messing with querying for a single record)
	sqlStatement := `
	SELECT id, email FROM users WHERE id=$1;
	`
	var email string
	var id int
	row := db.QueryRow(sqlStatement, 17)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Println("ID and email are:")
		fmt.Println(id, email)
	default:
		panic(err)
	}
	sqlStatement = `
	SELECT * FROM users WHERE id=$1
	`
	var user User
	row = db.QueryRow(sqlStatement, 17)
	err = row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Println("User's full data is:")
		fmt.Println(user)
	default:
		panic(err)
	}
}
