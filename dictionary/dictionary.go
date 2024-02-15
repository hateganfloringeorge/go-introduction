package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
	ErrUpdateFailed = DictionaryErr("provided word does not have an entry to update")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrUpdateFailed
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}