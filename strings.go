package nautilus

import (
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
)

// ToPascal convert given string case to pascal case
// For example, if str is like "example_string" or
// "exampleString" the output will be ExampleString
func ToPascal(str string) string {
	return camelCaseGenerator(str, true)
}

// ToCamel convert given string case to camel case
// For example, if str is like "example_string" or
// "ExampleString" the output will be exampleString
func ToCamel(str string) string {
	return camelCaseGenerator(str, false)
}

// ToSnake convert given string case to snake case
// For example, if str is like "exampleString" or
// "ExampleString" the output will be example_string
func ToSnake(str string) string {
	return snakeCaseGenerator(str, '_', false)
}

// ToScreamingSnake convert given string case to all
// capital snake case.
// For example, if str is like "exampleString" or
// "ExampleString" the output will be EXAMPLE_STRING
func ToScreamingSnake(str string) string {
	return snakeCaseGenerator(str, '_', true)
}

// ToKebab is similar to ToSnake but instead of using
// underscore, it uses dash.
// For example, if str is like "exampleString" or
// "ExampleString" the output will be example-string
func ToKebab(str string) string {
	return snakeCaseGenerator(str, '-', false)
}

// ToScreamingKebab is similar to ToScreamingSnake but
// instead of using underscore, it uses dash.
// For example, if str is like "exampleString" or
// "ExampleString" the output will be EXAMPLE-STRING
func ToScreamingKebab(str string) string {
	return snakeCaseGenerator(str, '-', true)
}

// Plural return the plural version of given string
// It uses inflection package
func Plural(str string) string {
	return inflection.Plural(str)
}

// Singular return singular version of given string
// It uses inflection package
func Singular(str string) string {
	return inflection.Singular(str)
}

// snakeCaseGenerator generate snake or kebab case from given string
func snakeCaseGenerator(str string, delimiter uint8, screaming bool) string {
	str = addWordBoundariesToNumbers(str)
	str = strings.Trim(str, " ")

	result := ""

	for i, char := range str {
		nextCaseIsChanged := false

		if i+1 < len(str) {
			next := str[i+1]
			if (char >= 'A' && char <= 'Z' && next >= 'a' && next <= 'z') || (char >= 'a' && char <= 'z' && next >= 'A' && next <= 'Z') {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && result[len(result)-1] != delimiter && nextCaseIsChanged {
			if char >= 'A' && char <= 'Z' {
				result += string(delimiter) + string(char)
			} else if char >= 'a' && char <= 'z' {
				result += string(char) + string(delimiter)
			}
		} else if char == ' ' || char == '_' || char == '-' {
			result += string(delimiter)
		} else {
			result = result + string(char)
		}
	}

	if screaming {
		result = strings.ToUpper(result)
	} else {
		result = strings.ToLower(result)
	}
	return result
}

// camelCaseGenerator generate camel or pascal case from given string
func camelCaseGenerator(str string, upperInit bool) string {
	str = addWordBoundariesToNumbers(str)
	str = addWordBoundaries(str)
	str = strings.Trim(str, " ")
	result := ""

	capNext := upperInit

	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			if !capNext {
				result += strings.ToLower(string(char))
			} else {
				result += string(char)
			}
		} else if char >= '0' && char <= '9' {
			result += string(char)
		} else if char >= 'a' && char <= 'z' {
			if capNext {
				result += strings.ToUpper(string(char))
			} else {
				result += string(char)
			}
		}

		if char == '_' || char == ' ' || char == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}

	return result
}

// addWordBoundariesToNumbers add boundaries for words and numbers
func addWordBoundariesToNumbers(s string) string {
	num := regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	numberReplacement := `$1 $2 $3`
	return num.ReplaceAllString(s, numberReplacement)
}

func addWordBoundaries(s string) string {
	words := regexp.MustCompile(`([A-Z][a-z]+)([A-Z][a-z]+)|([A-Z]+)([A-Z][a-z]+)|([A-Z]+)`)
	numberReplacement := `$1 $2 $3 $4 $5`
	s = words.ReplaceAllString(s, numberReplacement)

	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, " ")
}
