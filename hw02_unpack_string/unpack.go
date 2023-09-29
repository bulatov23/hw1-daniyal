package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func validate(str string) bool { // здесь я написал функцию которая валидирует строку на некорректные значени, такие как цифра в первом символе или число больше 9)
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

func Unpack(str string) (string, error) { // здесь получаю строку, вызываю функцию валидации и распаковываю строку
	runes := []rune(str)
	if str == "" {
		return "", nil
	}
	val := validate(str)
	if val == false {
		return "некорректная строка", nil
	}
	//var builder strings.Builder
	previousChar := ""
	for ind, char := range runes {
		if unicode.IsDigit(char) {
			count, _ := strconv.Atoi(string(char))
			if count == 0 {
				previousChar = string(runes[:ind-1])
			} else if count > 0 {
				if string(str[ind-1]) == "\n" {
					previousChar += strings.Repeat("\\n", count-1)
				} else {
					previousChar += strings.Repeat(string(runes[ind-1]), count-1)
				}
			}
		} else if !unicode.IsDigit(char) {
			if string(char) == "\n" {
				previousChar += "\\n"
			} else {
				previousChar += string(char)
			}
		}
	}
	return previousChar, nil
}
