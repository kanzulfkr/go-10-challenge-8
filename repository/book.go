package repository

import (
	"C8/model"
)

type BookRepo interface {
	CreateBook(model.Book) (res model.Book, err error)
	ListBook() (res []model.Book, err error)
	GetBook(id string) (res model.Book, err error)
	UpdateBook(id string, newBook model.Book) (res model.Book, err error)
	DeleteBook(id string) (count int64, err error)
}

func (r Repo) CreateBook(newBook model.Book) (res model.Book, err error) {
	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *
	`

	err = r.db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Description).Scan(&res.ID, &res.Title, &res.Author, &res.Description)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) ListBook() (res []model.Book, err error) {
	sqlStatement := `
	SELECT * FROM books ORDER BY id
	`

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp model.Book

		err = rows.Scan(&temp.ID, &temp.Title, &temp.Author, &temp.Description)

		if err != nil {
			return res, err
		}

		res = append(res, temp)
	}

	return res, nil
}

func (r Repo) GetBook(id string) (res model.Book, err error) {
	sqlStatement := `
	SELECT * FROM books WHERE id = $1
	`

	rows, err := r.db.Query(sqlStatement, id)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&res.ID, &res.Title, &res.Author, &res.Description)

		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r Repo) UpdateBook(id string, newBook model.Book) (res model.Book, err error) {
	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1
	Returning *
	`

	err = r.db.QueryRow(sqlStatement, id, newBook.Title, newBook.Author, newBook.Description).Scan(&res.ID, &res.Title, &res.Author, &res.Description)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id string) (count int64, err error) {
	sqlStatement := `
	DELETE from books
	WHERE id = $1
	`

	res, err := r.db.Exec(sqlStatement, id)

	if err != nil {
		return 0, err
	}

	count, err = res.RowsAffected()
	if err != nil {
		return count, err
	}

	return count, nil
}
