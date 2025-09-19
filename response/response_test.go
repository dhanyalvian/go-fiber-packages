//- response/response_test.go

package response

import (
	"net/http"
	"testing"

	"github.com/dhanyalvian/go-fiber-packages/base"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetResponseReqId(t *testing.T) {
	app := fiber.New()

	// Acquire fiber.Ctx from fasthttp.RequestCtx
	ctx := base.GetCtx(app)
	defer app.ReleaseCtx(ctx)

	// Test when X-Request-ID header is not set
	reqID := GetResponseReqId(ctx)
	assert.Equal(t, "", reqID)

	// Set X-Request-ID header
	expectedID := "test-req-id-123"
	ctx.Response().Header.Set(fiber.HeaderXRequestID, expectedID)

	// Test when X-Request-ID header is set
	reqID = GetResponseReqId(ctx)
	assert.Equal(t, expectedID, reqID)
}

func TestGetResponseStatusCode(t *testing.T) {
	app := fiber.New()

	// Acquire fiber.Ctx from fasthttp.RequestCtx
	ctx := base.GetCtx(app)
	defer app.ReleaseCtx(ctx)

	// Set status code
	ctx.Response().SetStatusCode(http.StatusAccepted) // 202

	// Call the function under test
	code := GetResponseStatusCode(ctx)

	// Assert
	assert.Equal(t, http.StatusAccepted, code)
}
