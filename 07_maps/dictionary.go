package main

//import "errors"

type Dictionary map[string]string
type DictionaryErr string
const(
	ErrNotFound = DictionaryErr("could not find the word you are looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)
func (e DictionaryErr) Error() string{
	return string(e)
}


func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	// word not found -> err = ErrNotFound -> go ahead add
	// word found -> err = nil -> return ErrWordExists
	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}