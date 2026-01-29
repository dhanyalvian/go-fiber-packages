//- generators/nanoid.go

package generators

import nanoid "github.com/matoous/go-nanoid/v2"

func NanoID(length int, uppercase bool) string {
	seed := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if !uppercase {
		seed += "abcdefghijklmnopqrstuvwxyz"
	}

	result, _ := nanoid.Generate(seed, length)
	return result
}
