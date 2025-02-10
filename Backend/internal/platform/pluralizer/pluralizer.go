// Package pluralizer is the wrapper for the go-pluralize package
package pluralizer

import "github.com/gertd/go-pluralize"

// ToPlural is the wrapper for the pluralize Plural function
func ToPlural(word string) string {
	pluralize := pluralize.NewClient()

	return pluralize.Plural(word)
}
