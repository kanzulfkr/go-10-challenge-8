package service

import "C8/model"

type BookService interface {
	CreateBook(model.Book) (res model.Book, err error)
	ListBook() (res []model.Book, err error)
	GetBook(id string) (res model.Book, err error)
	UpdateBook(id string, newBook model.Book) (res model.Book, err error)
	DeleteBook(id string) (count int64, err error)
}

func (s *Service) CreateBook(newBook model.Book) (res model.Book, err error) {
	res, err = s.repo.CreateBook(newBook)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) ListBook() (res []model.Book, err error) {
	res, err = s.repo.ListBook()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBook(id string) (res model.Book, err error) {
	res, err = s.repo.GetBook(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UpdateBook(id string, newBook model.Book) (res model.Book, err error) {
	res, err = s.repo.UpdateBook(id, newBook)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) DeleteBook(id string) (count int64, err error) {
	count, err = s.repo.DeleteBook(id)
	if err != nil {
		return count, err
	}
	return count, nil
}
