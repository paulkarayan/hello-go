package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("couldnt find word")

//thin wrapper around map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
