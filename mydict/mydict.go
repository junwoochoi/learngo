package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errCantUpdate = errors.New("Cant update non-existing word")
var errWordExists = errors.New("That word already exists")

//Search in dictionary
func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]

	if !exists {
		return "", errNotFound
	}
	return value, nil
}

//Add word with def
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}

	return nil
}

//Update word with def
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		return errCantUpdate
	} else if err == nil {
		d[word] = def
	}

	return nil
}

//Delete word in dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
