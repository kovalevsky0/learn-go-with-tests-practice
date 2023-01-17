package dictionary

type Dictionary map[string]string

// Constant Errors: https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryError string

var (
	ErrNotFound         = DictionaryError("There is no such word in the dictionary")
	ErrWordExists       = DictionaryError("This word is already in the dictionary")
	ErrWordDoesNotExist = DictionaryError("Cannot update/delete the word because it does not exist in the dictionary")
)

func (d DictionaryError) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

// "An interesting property of maps is that you can modify them without passing as an address to it"
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[word] = definition
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
		return ErrWordDoesNotExist
	default:
		return err
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}
