package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

/*
Names holds a slice of names that will be displayed on the webpage
*/
type Names struct {
	NamesList []Name
}

/*
Name is a structure with two strings
One for the first name and one for the last name
*/
type Name struct {
	ID    int
	FName string
	LName string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//dbConn is used to make a connection with the database
func dbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:Enter123$@tcp(mydb:3306)/testdb")
	checkError(err)
	return db
}

func getAllNames(db *sql.DB) (*Names, error) {
	//create a new slice to hold data
	var n Names

	rows, err := db.Query("SELECT * FROM names ORDER BY fName ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		//assign value to strings
		var idString, fName, lName string
		if err = rows.Scan(&idString, &fName, &lName); err != nil {
			log.Fatal(err)
		}

		id, _ := strconv.Atoi(idString)

		//create a Name struct using input
		name := Name{FName: fName, LName: lName, ID: id}

		//add new Name struct to slice
		n.NamesList = append(n.NamesList, name)

	}

	return &n, nil
}

//each handler will process and display tha proper info in the webpage
func viewHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	//pull the data from the database
	names, err := getAllNames(db)
	checkError(err)

	t, _ := template.ParseFiles("./index.html")

	t.Execute(w, names)
}

func addNameHandler(w http.ResponseWriter, r *http.Request) {
	//connect to database
	db := dbConn()
	defer db.Close()

	//read the form data
	err := r.ParseForm()
	checkError(err)

	//assign value from form to a variable
	newFName := r.Form.Get("fName")
	newLName := r.Form.Get("lName")

	//add new entry into database
	sql := "INSERT INTO names (fName, lName) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	checkError(err)
	defer stmt.Close()
	_, err = stmt.Exec(newFName, newLName)
	checkError(err)

	//pull the updated data from the database
	names, err := getAllNames(db)
	checkError(err)

	//reload the webpage to reflect the changes
	t, _ := template.ParseFiles("./index.html")

	t.Execute(w, names)
}

func deleteNameHandler(w http.ResponseWriter, r *http.Request) {
	//connect to database
	db := dbConn()
	defer db.Close()

	testString := ""
	fmt.Println(r.URL.Parse(testString))

	//read data to be removed
	input := strings.Split(r.URL.Path, "/")

	fmt.Println(input)

	//var name Name
	fullName := strings.Split(input[2], " ")

	fmt.Println(fullName)

	//delete entry from database
	sql := "DELETE FROM names WHERE (fName) IN (?)"
	stmt, err := db.Prepare(sql)
	checkError(err)
	defer stmt.Close()
	_, err = stmt.Exec(strings.Replace(fullName[0], "{", "", 1))
	checkError(err)

	//pull the updated data from the database
	names, err := getAllNames(db)
	checkError(err)

	//reload the webpage to reflect the changes
	t, _ := template.ParseFiles("./index.html")

	t.Execute(w, names)

}

func updateNameHandler(w http.ResponseWriter, r *http.Request) {
	//connect to database
	db := dbConn()
	defer db.Close()

	//read data to be updated
	err := r.ParseForm()
	checkError(err)

	//assign value from form to a variable
	id := r.Form.Get("id")
	newFName := r.Form.Get("updateFName")
	newLName := r.Form.Get("updateLName")

	//update entry in the database
	sql := "UPDATE names SET fNAME=?, lName=? WHERE name_id =?"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	_, err = stmt.Exec(newFName, newLName, id)
	checkError(err)

	//pull the updated data from the database
	names, err := getAllNames(db)
	checkError(err)

	//reload the webpage to reflect the changes
	t, err := template.ParseFiles("./index.html")
	checkError(err)

	t.Execute(w, names)

}

func main() {

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/addName/", addNameHandler)
	http.HandleFunc("/deleteName/", deleteNameHandler)
	http.HandleFunc("/updateName/", updateNameHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
