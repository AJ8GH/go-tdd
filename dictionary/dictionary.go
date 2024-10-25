package dictionary

type Dictionary map[string]string

var (
	ErrNoWordFound error = DictionaryErr{"no word found"}
	ErrWordExists  error = DictionaryErr{"word exists"}
)

type DictionaryErr struct {
	message string
}

func (e DictionaryErr) Error() string {
	return e.message
}

func (d Dictionary) Search(word string) (string, error) {
	got, exists := d[word]
	if !exists {
		return "", ErrNoWordFound
	}
	return got, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, exists := d[word]
	if exists {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}
