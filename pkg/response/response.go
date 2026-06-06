package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, R{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, R{
		Code:    201,
		Message: "created",
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, R{
		Code:    code,
		Message: msg,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, R{
		Code:    400,
		Message: msg,
	})
}

func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, R{
		Code:    401,
		Message: msg,
	})
	c.Abort()
}

func Forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, R{
		Code:    403,
		Message: msg,
	})
	c.Abort()
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, R{
		Code:    404,
		Message: msg,
	})
}

func InternalError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, R{
		Code:    500,
		Message: msg,
	})
}

// Error codes
const (
	ErrUserNotFound      = 10001
	ErrPasswordIncorrect = 10002
	ErrPhoneExists       = 10003
	ErrUserDisabled      = 10004

	ErrCourseNotFound = 20001
	ErrCoursePurchased = 20002
	ErrChapterNotFound = 20003

	ErrVideoNotFound = 30001
	ErrVideoNotReady = 30002

	ErrExamNotFound   = 40001
	ErrExamSubmitted  = 40002
	ErrQuestionFailed = 40003

	ErrOrderExists    = 50001
	ErrOrderNotFound  = 50002
	ErrPaymentFailed  = 50003
)
