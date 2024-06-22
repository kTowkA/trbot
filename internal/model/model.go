package model

type AuthorDictionary struct {
	Name       string
	Additional string
}
type Language struct {
	Title      string
	Additional string
}
type Dictionary struct {
	Title        string
	DateCreating string
	Additional   string
	Languages    []Language
	Authors      []AuthorDictionary
}

type TranslateElement struct {
	Value                     string
	ValueWithAccent           string
	TranslatedValue           string
	TranslatedValueWithAccent string
	SourceLanguage            Language
	TranslationLanguage       Language
	Dictionary                Dictionary
}

type Word TranslateElement
type Phrase TranslateElement
