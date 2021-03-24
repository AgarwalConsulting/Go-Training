package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

var hello = "hello" // 1. Initialize package levels

func init() {
	log.Println("Init from main...")
}

type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"title"`
	ProjectID   int    `json:"project" validate:"gte=1000"`
}

func main() {
	// db, err := sql.Open("sqlite", "file:/tmp/employee.sqlite")
	db, err := sql.Open("postgres", "postgres://localhost:5432/employees?sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Println("Error:", err)
		return
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, name TEXT, project_id INTEGER, designation TEXT);")

	if err != nil {
		log.Println("Unable to create table:", err)
		return
	}

	rows, err := db.Query("SELECT * FROM employees WHERE name = $1", "Gaurav")

	if err != nil {
		log.Println("Unable to select:", err)
		return
	}

	if rows.Next() {
		var emp Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.ProjectID, &emp.Designation)

		log.Println(emp)
	} else {
		log.Println("No rows found!")
	}

	log.Println("Preparing statement...")
	stmt, err := db.Prepare("INSERT INTO employees (id, name, project_id, designation) VALUES($1, $2, $3, $4)")
	defer stmt.Close()

	if err != nil {
		log.Fatal("Unable to prepare statement:", err)
	}

	log.Println("Inserting...")
	// res, err := db.Exec("INSERT INTO employees (id, name, project_id, designation) VALUES($1, $2, $3, $4)", 2, "Kiran", 1002, "Intern")
	// res, err := stmt.Exec(2, "Kiran", 1002, "Intern")
	res, err := stmt.Exec(1, "Gaurav", 1001, "Developer")

	if err != nil {
		log.Println("Unable to insert:", err)
		return
	}

	fmt.Println("Inserted successfully...")

	rowsAffected, _ := res.RowsAffected()
	log.Println("No. of rows affected:", rowsAffected)
}
