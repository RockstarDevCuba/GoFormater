package main

import (
	"fmt"
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

func Parser(str string, _map MapArgv) (string, error) {
	mc := regexp.MustCompile(parserPattern)
	allFormatParser := mc.FindAllString(str, -1)
	for i := range allFormatParser {
		_validation := validate(allFormatParser[i])
		if !_validation {
			return str, err.InvalidFormatError(allFormatParser[i])
		}
		replace, err := getElementReplace(allFormatParser[i], _map)
		fmt.Println(replace)
		if err != nil {
			return str, err
		}
		str = strings.Replace(str, allFormatParser[i], replace, -1)
	}
	return str, nil
}
