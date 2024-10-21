package dictionary

import "errors"

type Dictionary map[string]string

var ErrNoWordFound = errors.New("no word found")

func (d Dictionary) Search(word string) (string, error) {
	got, exists := d[word]
	if !exists {
		return "", ErrNoWordFound
	}
	return got, nil
}
