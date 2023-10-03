package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func validate(str string) bool {
	strArr := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, s := range strArr {
		if string(str[0]) == s {
			return false
		}
	}
	re := regexp.MustCompile("[0-9]+")
	arr := re.FindAllString(str, -1)
	for _, v := range arr {
		num, _ := strconv.Atoi(v)
		if num > 9 {
			return false
		}
	}
	return true
}

func Unpack(str string) (string, error) {
	var builder strings.Builder
	runes := []rune(str)
	if str == "" {
		return "", nil
	}
	val := validate(str)
	if !val {
		return "", ErrInvalidString
	}
	for ind, char := range runes {
		if ind == len(runes)-1 {
			continue
		}
		count, _ := strconv.Atoi(string(runes[ind+1]))
		if count == 0 && unicode.IsDigit(char) {
			continue
		}
		if unicode.IsDigit(runes[ind+1]) {
			builder.WriteString(strings.Repeat(string(runes[ind]), count))
		} else if !unicode.IsDigit(runes[ind+1]) {
			builder.WriteString(string(runes[ind]))
		}
	}
	return builder.String(), nil
}
