//- utils/generator.go

package utils

import nanoid "github.com/matoous/go-nanoid/v2"

func NanoID(length int) string {
	result, _ := nanoid.Generate(
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		length,
	)
	return result
}
