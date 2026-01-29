//- response/response.go

package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Meta    ResponseMeta `json:"meta"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseMeta struct {
	RequestId string `json:"reqId"`
	Code      string `json:"code"`
}

type ResponseData struct {
	Pagination *ResponseDataPagination `json:"pagination,omitempty"`
	Results    interface{}             `json:"results,omitempty"`
	Result     interface{}             `json:"result,omitempty"`
	Errors     interface{}             `json:"errors,omitempty"`
}

type ResponseDataPagination struct {
	Page         int   `json:"page"`
	Next         int   `json:"next"`
	Records      int   `json:"records"`
	TotalPages   int   `json:"totalPages"`
	TotalRecords int64 `json:"totalRecords"`
}

type ResponseLog struct {
	ReqId      string      `json:"reqId,omitempty"`
	Headers    interface{} `json:"headers,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

func GetResponseReqId(c *fiber.Ctx) string {
	return string(c.Response().Header.Peek(fiber.HeaderXRequestID))
}

func GetResponseStatusCode(c *fiber.Ctx) int {
	return c.Response().StatusCode()
}
