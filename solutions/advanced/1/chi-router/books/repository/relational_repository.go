package repository

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3" // It call init() within the package to register the sqlite3

	"biblioteca/entities"
)

type relationalRepo struct {
	*sql.DB
}

func (repo relationalRepo) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	rows, err := repo.DB.Query("SELECT id, title, author, price FROM books")

	if err != nil {
		log.Println("Unable to retrieve from books:", err)
		return nil, err
	}

	for rows.Next() {
		var book entities.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)

		if err != nil {
			log.Println("Unable to scan:", err)
			return nil, err
		}

		log.Printf("Scanned: %#v\n", book)

		log.Println("\tId:", book.ID)
		log.Println("\tName:", book.Title)
		log.Println("\tAuthor:", book.Author)
		// log.Println("\tIsbn:", *book.isbn)
		log.Println("\tPrice:", book.Price)
		// log.Println("\tDescription:", *book.description)

		books = append(books, book)
	}

	return books, nil
}

func (relationalRepo) FindBy(id int) (entities.Book, error) {
	return entities.Book{}, nil
}

func (relationalRepo) Create(newBook entities.Book) (*entities.Book, error) {
	return nil, nil
}

func (relationalRepo) Update(existingBook *entities.Book) error {
	return nil
}

func (relationalRepo) Destroy(existingBook *entities.Book) error {
	return nil
}

func NewRelationRepository() BookRepository {
	db, err := sql.Open("sqlite", "./books.db")

	if err != nil {
		log.Fatalln("Unable to open DB:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id numeric primary key, title text, author text, isbn text, price numeric, description text)")

	if err != nil {
		log.Fatalln("Unable to create table:", err)
	}

	return relationalRepo{db}
}
