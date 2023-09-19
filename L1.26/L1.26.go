package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, что все символы в строке уникальные (регистронезависимо)
func IsUnique(str string) bool {
	// Преобразуем строку к нижнему регистру для регистронезависимой проверки
	str = strings.ToLower(str)

	charCount := make(map[rune]int)

	// Подсчитываем количество вхождений каждого символа
	for _, char := range str {
		charCount[char]++
		if charCount[char] > 1 {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		str      string
		isUnique bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
	}

	for _, testCase := range testCases {
		result := IsUnique(testCase.str)
		fmt.Printf("Строка '%s' уникальная: %v\n", testCase.str, result)
	}
}
