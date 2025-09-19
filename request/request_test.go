package request

import (
	"testing"

	"github.com/dhanyalvian/go-fiber-packages/base"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		query    string
		expected int
	}{
		{"", 1},
		{"p=0", 1},
		{"p=-5", 1},
		{"p=2", 2},
		{"p=100", 100},
		{"p=abc", 1},
	}

	for _, tt := range tests {
		ctx := base.GetCtx(app)
		defer app.ReleaseCtx(ctx)

		ctx.Request().URI().SetQueryString(tt.query)
		assert.Equal(t, tt.expected, GetPage(ctx))
		app.ReleaseCtx(ctx)
	}
}

func TestGetLimit(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		query    string
		expected int
	}{
		{"", 20},
		{"l=0", 20},
		{"l=-10", 20},
		{"l=10", 10},
		{"l=114", 114},
		{"l=200", 114},
		{"l=abc", 20},
	}

	for _, tt := range tests {
		ctx := base.GetCtx(app)
		defer app.ReleaseCtx(ctx)

		ctx.Request().URI().SetQueryString(tt.query)
		assert.Equal(t, tt.expected, GetLimit(ctx))
		app.ReleaseCtx(ctx)
	}
}

func TestCountItems(t *testing.T) {
	assert.Equal(t, 0, CountItems(nil))
	assert.Equal(t, 0, CountItems(123))
	assert.Equal(t, 0, CountItems(struct{}{}))
	assert.Equal(t, 3, CountItems([]int{1, 2, 3}))
	assert.Equal(t, 2, CountItems([2]string{"a", "b"}))
	assert.Equal(t, 0, CountItems([]string{}))
}
