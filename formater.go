package goformater

import (
	"regexp"
	"strings"

	err "github.com/RockstarDevCuba/GoFormater/extra"
)

type MapArgv map[string]string

const parserPattern = `{.*?}`

var oldnew = []string{"{", "", "}", ""}

func validate(args string) bool {
	return !strings.Contains(args, " ")
}

func getElementReplace(keys string, _map MapArgv) (string, error) {
	key := strings.NewReplacer(oldnew...).Replace(keys)
	val, isExist := _map[key]
	if !isExist {
		return "", err.KeyError(key)
	}
	return val, nil
}

// Parser receives a text string and an argument of type MapArgv equals map[string]string
// of type MapArgv equals map[string]string
// Returns strin
func Parser(str string, _map MapArgv) (string, error) {
	mc := regexp.MustCompile(parserPattern)
	allFormatParser := mc.FindAllString(str, -1)
	for i := range allFormatParser {
		_validation := validate(allFormatParser[i])
		if !_validation {
			return str, err.InvalidFormatError(allFormatParser[i])
		}
		replace, err := getElementReplace(allFormatParser[i], _map)
		if err != nil {
			return str, err
		}
		str = strings.Replace(str, allFormatParser[i], replace, -1)
	}
	return str, nil
}

// ParseByArgs takes a text string,
// true/false to throw errors and arguments in the form of an array
// return string
// If errors = true it will throw an exception of type IndexError
func ParserByArgs(str string, errors bool, args ...string) (string, error) {
	var index int16
	mc := regexp.MustCompile("{}")
	allFormat := mc.FindAllString(str, -1)
	index = int16(len(allFormat))
	if errors && (int16(len(args)) < index) {
		return str, err.IndexError("The number of arguments received cannot be less than the number of \"{}\" characters in the text string")
	}
	for i := range allFormat {
		if i > len(args)-1 {
			break
		}
		str = strings.Replace(str, allFormat[i], args[i], 1)
	}
	return str, nil
}
