//This is in its own branch just to test how to make a branch but yeah

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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//This will update Jake Baldino and
	//also will spit out the ID and email of the mentioned row
	sqlStatement := `
	UPDATE users
	SET first_name = $2, last_name = $3
	WHERE id = $1
	RETURNING id, email;
	`
	var email string
	var id int
	err = db.QueryRow(sqlStatement, 3, "Fake", "Baldino").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID and email are:")
	fmt.Println(id, email)

	//Deleting Pappa's Yahoo account >:)
	sqlStatement = `
	DELETE FROM users
	WHERE id = $1;
	`
	_, err = db.Exec(sqlStatement, 39)
	if err != nil {
		panic(err)
	}

}
