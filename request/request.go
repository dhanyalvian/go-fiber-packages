// - request/request.go

package request

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPage(c *fiber.Ctx) int {
	page, _ := strconv.Atoi(c.Query("p"))
	if page <= 0 {
		page = 1
	}

	return page
}

func GetLimit(c *fiber.Ctx) int {
	limit, _ := strconv.Atoi(c.Query("l"))
	switch {
	case limit > 114:
		limit = 114
	case limit <= 0:
		limit = 20
	}

	return limit
}

func CountItems(v interface{}) int {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		return rv.Len()
	}

	return 0
}
