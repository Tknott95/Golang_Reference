package cms

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var store = newDB()

type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=postgres dbname=gocmsdb2 sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}

func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}

func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}
