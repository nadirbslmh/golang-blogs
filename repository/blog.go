package repository

import (
	"golang-blogs/database"
	"golang-blogs/model"
)

type BlogRepositoryImpl struct{}

func (b *BlogRepositoryImpl) Create(input model.BlogInput) {
	stmt, err := database.DB.Prepare("INSERT INTO blog(title,content) VALUES(?,?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(input.Title, input.Content)
	if err != nil {
		panic(err)
	}
}
