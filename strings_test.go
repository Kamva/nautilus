package nautilus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strs = []string{
	"ExampleString",
	"exampleString",
	"example_string",
	"Example_String",
	"EXAMPLE_STRING",
	"example string",
	"Example string",
	"Example String",
	"EXAMPLE STRING",
	"ExampleString2",
	"exampleString2",
	"example_string2",
	"example_string_2",
	"Example_String2",
	"Example_String_2",
	"EXAMPLE_STRING2",
	"EXAMPLE_STRING_2",
	"example string 2",
	"example string2",
	"Example string 2",
	"Example string2",
	"Example String 2",
	"Example String2",
	"EXAMPLE STRING 2",
	"EXAMPLE STRING2",
}

var pascalResult = "ExampleString"
var pascalResultWithNum = "ExampleString2"
var camelResult = "exampleString"
var camelResultWithNum = "exampleString2"
var snakeResult = "example_string"
var snakeResultWithNum = "example_string_2"
var screamSnakeResult = "EXAMPLE_STRING"
var screamSnakeResultWithNum = "EXAMPLE_STRING_2"
var kebabResult = "example-string"
var kebabResultWithNum = "example-string-2"
var screamKebabResult = "EXAMPLE-STRING"
var screamKebabResultWithNum = "EXAMPLE-STRING-2"

var words = map[string]string{
	"star":        "stars",
	"bus":         "buses",
	"mouse":       "mice",
	"query":       "queries",
	"agency":      "agencies",
	"movie":       "movies",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"datum":       "data",
	"medium":      "media",
	"analysis":    "analyses",
	"child":       "children",
	"news":        "news",
	"series":      "series",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
}

func TestToPascal(t *testing.T) {
	for i, str := range strs {
		rs := ToPascal(str)

		if i < 9 {
			assert.Equal(t, pascalResult, rs)
		} else {
			assert.Equal(t, pascalResultWithNum, rs)
		}
	}
}

func TestToCamel(t *testing.T) {
	for i, str := range strs {
		rs := ToCamel(str)

		if i < 9 {
			assert.Equal(t, camelResult, rs)
		} else {
			assert.Equal(t, camelResultWithNum, rs)
		}
	}
}

func TestToSnake(t *testing.T) {
	for i, str := range strs {
		rs := ToSnake(str)

		if i < 9 {
			assert.Equal(t, snakeResult, rs)
		} else {
			assert.Equal(t, snakeResultWithNum, rs)
		}
	}
}

func TestToScreamingSnake(t *testing.T) {
	for i, str := range strs {
		rs := ToScreamingSnake(str)

		if i < 9 {
			assert.Equal(t, screamSnakeResult, rs)
		} else {
			assert.Equal(t, screamSnakeResultWithNum, rs)
		}
	}
}

func TestToKebab(t *testing.T) {
	for i, str := range strs {
		rs := ToKebab(str)

		if i < 9 {
			assert.Equal(t, kebabResult, rs)
		} else {
			assert.Equal(t, kebabResultWithNum, rs)
		}
	}
}

func TestToScreamingKebab(t *testing.T) {
	for i, str := range strs {
		rs := ToScreamingKebab(str)

		if i < 9 {
			assert.Equal(t, screamKebabResult, rs)
		} else {
			assert.Equal(t, screamKebabResultWithNum, rs)
		}
	}
}

func TestPlural(t *testing.T) {
	for singular, plural := range words {
		res := Plural(singular)
		assert.Equal(t, plural, res)
	}
}

func TestSingular(t *testing.T) {
	for singular, plural := range words {
		res := Singular(plural)
		assert.Equal(t, singular, res)
	}
}
