package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1:9042")

	cluster.Keyspace = "library"

	session, err := cluster.CreateSession()

	if err != nil {
		log.Println("Unable to create session:", err)
		return
	}
	defer session.Close()

	// keyspace, err := session.KeyspaceMetadata("library")

	// if err != nil {
	// 	log.Println("Unable to get keyspace:", err)
	// }

	// log.Println("Keyspace:", keyspace)

	// if err := session.Query("CREATE KEYSPACE IF NOT EXISTS library WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}").Exec(); err != nil {
	// 	log.Println("Unable to create keyspace:", err)
	// 	return
	// }

	// keyspace, _ = session.KeyspaceMetadata("library")

	// log.Println("Keyspace:", keyspace)

	// if err := session.Query("USE library").Exec(); err != nil {
	// 	log.Println("Unable to use library keyspace:", err)
	// 	return
	// }

	if err := session.Query("CREATE TABLE IF NOT EXISTS books (id INT PRIMARY KEY, title TEXT, author TEXT, isbn TEXT, description TEXT, price double)").Exec(); err != nil {
		log.Println("Unable to create table books:", err)
		return
	}

	// err = session.Query(`INSERT INTO books (id, title, author, description) VALUES (?, ?, ?, ?)`, 1, "One up on Wall Street", "Peter Lynch", "An amazing stock market investors guide").Exec()
	err = session.Query(`INSERT INTO books (id, title, author, description) VALUES (?, ?, ?, ?)`, 2, "Value investing & behavior finance", "Parag Parikh", "Psychology of avg stock investors").Exec()

	if err != nil {
		log.Println("Unable to insert into books:", err)
		return
	}

	var book struct {
		title  string
		author string
	}

	err = session.Query(`SELECT title, author FROM books`).Scan(&book.title, &book.author)

	if err != nil {
		log.Println("Unable to retrieve a book:", err)
		return
	}

	log.Println("Retrieved Book:", book)

	scanner := session.Query(`SELECT title, author FROM books`).Iter().Scanner()

	log.Println("Retrieving all books...")

	for scanner.Next() {
		scanner.Scan(&book.title, &book.author)

		log.Println("\tBook:", book)
	}
}
