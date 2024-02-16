package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (dict Dictionary) Search(key string) (string, error) {
	definition, ok := dict[key]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrorNotFound:
		dict[key] = value
	case nil:
	default:
		return err
	}
	return nil
}
