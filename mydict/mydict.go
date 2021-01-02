package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

//Search in dictionary
func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]

	if !exists {
		return "", errNotFound
	}
	return value, nil
}
