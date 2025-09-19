//- response/response.go

package response

import "github.com/gofiber/fiber/v2"

func GetResponseReqId(c *fiber.Ctx) string {
	return string(c.Response().Header.Peek(fiber.HeaderXRequestID))
}

func GetResponseStatusCode(c *fiber.Ctx) int {
	return c.Response().StatusCode()
}
