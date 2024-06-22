package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/kTowkA/trbot/internal/model"
)

type StorageAuthor interface {
	// SaveAuthor сохраняет нового автора словаря author. Возвращает UUID и ошибку. Может вернуть ErrAuthorAlreadyExists, если такой автор уже существует.
	SaveAuthor(ctx context.Context, author model.AuthorDictionary) (uuid.UUID, error)
	// AuthorByName получает автора по имени name. Может вернуть ErrAuthorNotFound если такой автор не будет найден
	AuthorByName(ctx context.Context, name string) (model.AuthorDictionary, error)
	// Author получает автора по ID. Может вернуть ErrAuthorNotFound если такой автор не будет найден
	Author(ctx context.Context, authorID uuid.UUID) (model.AuthorDictionary, error)
}
type StorageLanguage interface {
	// SaveLanguage сохраняет новый язык language. Возвращает UUID и ошибку. Может вернуть ErrLanguageAlreadyExists, если такой язык уже существует.
	SaveLanguage(ctx context.Context, language model.Language) (uuid.UUID, error)
	// LanguageByName получает язык по названию title. Может вернуть ErrLanguageNotFound если такой язык не будет найден
	LanguageByName(ctx context.Context, title string) (model.Language, error)
	// Language получает язык по ID. Может вернуть ErrLanguageNotFound если такой язык не будет найден
	Language(ctx context.Context, languageID uuid.UUID) (model.Language, error)
}
type StorageDictionary interface {
	// SaveDictionary сохраняет новый словарь dictionary. Возвращает UUID и ошибку. Может вернуть ErrDictionaryAlreadyExists, если такой словарь уже существует.
	SaveDictionary(ctx context.Context, dictionary model.Dictionary) (uuid.UUID, error)
	// DictionaryByName получает словарь по названию title. Может вернуть ErrDictionaryNotFound если такой словарь не будет найден
	DictionaryByName(ctx context.Context, title string) (model.Dictionary, error)
	// Dictionary получает словарь по ID. Может вернуть ErrDictionaryNotFound если такой словарь не будет найден
	Dictionary(ctx context.Context, languageID uuid.UUID) (model.Dictionary, error)
}
