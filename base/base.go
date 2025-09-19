package base

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func GetCtx(app *fiber.App) *fiber.Ctx {
	// Create a fasthttp.RequestCtx manually
	var req fasthttp.Request
	fctx := fasthttp.RequestCtx{}
	fctx.Init(&req, nil, nil)

	return app.AcquireCtx(&fctx)
}
