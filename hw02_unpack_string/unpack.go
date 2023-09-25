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
	if str == "" {
		return "", nil
	}
	val := validate(str)
	if !val {
		return "", ErrInvalidString
	}
	previousChar := ""
	for ind, char := range str {
		if unicode.IsDigit(char) {
			count, _ := strconv.Atoi(string(char))
			if count == 0 {
				previousChar = str[:ind-1]
			} else if count > 0 {
				if str[ind-1] == '\n' {
					previousChar += strings.Repeat("\\n", count-1)
				} else {
					previousChar += strings.Repeat(string(str[ind-1]), count-1)
				}
			}
		} else if !unicode.IsDigit(char) {
			if char == '\n' {
				previousChar += "\\n"
			} else {
				previousChar += string(char)
			}
		}
	}
	return previousChar, nil
}
