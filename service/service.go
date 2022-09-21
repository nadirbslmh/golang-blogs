package service

import (
	"golang-blogs/model"
	"golang-blogs/repository"
)

type BlogService struct {
	Repository repository.BlogRepository
}

func New() BlogService {
	return BlogService{
		Repository: &repository.BlogRepositoryImpl{},
	}
}

func (b *BlogService) Create(input model.BlogInput) {
	b.Repository.Create(input)
}
