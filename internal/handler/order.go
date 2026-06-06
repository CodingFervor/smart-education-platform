package handler

import (
	"github.com/CodingFervor/smart-education-platform/pkg/response"
	"github.com/gin-gonic/gin"
)

type CreateOrderReq struct {
	CourseID int64  `json:"course_id" binding:"required"`
	CouponID *int64 `json:"coupon_id"`
}

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

// CreateOrder godoc
// @Summary Purchase a course
// @Tags Order
// @Security BearerAuth
// @Param body body CreateOrderReq true "Order data"
// @Success 201
// @Router /api/v1/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	userID := c.GetInt64("user_id")
	// TODO: check if already purchased, create order, return payment URL
	response.Created(c, gin.H{
		"order_no":  "EDU20260606001",
		"user_id":   userID,
		"course_id": req.CourseID,
		"pay_url":   "https://pay.example.com/xxx",
	})
}

// GetOrder godoc
// @Summary Get order detail
// @Tags Order
// @Security BearerAuth
// @Param order_no path string true "Order number"
// @Success 200
// @Router /api/v1/orders/{order_no} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderNo := c.Param("order_no")
	response.OK(c, gin.H{"order_no": orderNo, "status": 1})
}

// ListMyOrders godoc
// @Summary List current user orders
// @Tags Order
// @Security BearerAuth
// @Success 200
// @Router /api/v1/orders [get]
func (h *OrderHandler) ListMyOrders(c *gin.Context) {
	// TODO: query user orders
	response.OK(c, gin.H{"records": []interface{}{}})
}

// PayCallback godoc
// @Summary Payment callback (wechat/alipay)
// @Tags Order
// @Success 200
// @Router /api/v1/pay/callback/wechat [post]
func (h *OrderHandler) PayCallback(c *gin.Context) {
	// TODO: verify signature, update order status, grant course access
	c.String(200, "<xml><return_code><![CDATA[SUCCESS]]></return_code></xml>")
}

// GetLearningProgress godoc
// @Summary Get learning progress for a course
// @Tags Learning
// @Security BearerAuth
// @Param course_id path int true "Course ID"
// @Success 200
// @Router /api/v1/learning/{course_id}/progress [get]
func (h *OrderHandler) GetLearningProgress(c *gin.Context) {
	// TODO: query progress
	response.OK(c, gin.H{
		"progress":  65,
		"completed": false,
		"chapters":  []interface{}{},
	})
}

// UpdateProgress godoc
// @Summary Update learning progress (video position)
// @Tags Learning
// @Security BearerAuth
// @Param course_id path int true "Course ID"
// @Param chapter_id query int true "Chapter ID"
// @Param position query int true "Video position in seconds"
// @Success 200
// @Router /api/v1/learning/{course_id}/progress [put]
func (h *OrderHandler) UpdateProgress(c *gin.Context) {
	// TODO: update progress, check completion, issue certificate if 100%
	response.OK(c, nil)
}
