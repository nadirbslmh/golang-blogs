package repository

import "golang-blogs/model"

type BlogRepository interface {
	Create(input model.BlogInput)
}
