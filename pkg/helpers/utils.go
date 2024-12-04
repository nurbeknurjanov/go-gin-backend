package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func EncryptString(password string) (string, error) {
	encryptedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(encryptedBytes), nil
}

func Ucfirst(str string) string {
	caser := cases.Title(language.Und, cases.NoLower)
	return caser.String(str)
}

func FirstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

func FirstToUpper(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToUpper(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func SetField(obj interface{}, fieldName string, value interface{}) {
	oe := reflect.ValueOf(obj).Elem()
	fv := oe.FieldByName(fieldName)

	t := fv.Type()

	if fv.IsValid() && fv.CanSet() {
		fv.Set(reflect.ValueOf(value).Convert(t))
	}
}
