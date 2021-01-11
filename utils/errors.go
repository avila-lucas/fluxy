package utils

import "errors"

var FileNotFound = errors.New("file not found")
var ParsingError = errors.New("parsing error")

func IsFileNotFound(err error) bool {
	return err == FileNotFound
}

func IsParsingError(err error) bool {
	return err == ParsingError
}
