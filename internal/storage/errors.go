package storage

import "errors"

var (
	ErrAuthorAlreadyExists = errors.New("такой автор уже существует")
	ErrAuthorNotFound      = errors.New("такой автор не существует")

	ErrLanguageAlreadyExists = errors.New("такой язык уже существует")
	ErrLanguageNotFound      = errors.New("такой язык не существует")

	ErrDictionaryAlreadyExists = errors.New("такой словарь уже существует")
	ErrDictionaryNotFound      = errors.New("такой словарь не существует")
)
