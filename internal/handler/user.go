package handler

import (
	"strconv"

	"github.com/CodingFervor/smart-education-platform/internal/model"
	"github.com/CodingFervor/smart-education-platform/pkg/response"
	"github.com/gin-gonic/gin"
)

// ==================== DTOs ====================

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Role     string `json:"role" binding:"required,oneof=student teacher"` // admin only via backend
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *model.User `json:"user"`
}

type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// ==================== Handlers ====================

type UserHandler struct {
	// userService UserService (to be injected)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register godoc
// @Summary Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param body body RegisterReq true "Registration data"
// @Success 201 {object} model.User
// @Router /api/v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// TODO: call service layer
	// user, err := h.userService.Register(c.Request.Context(), &req)

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Role:     req.Role,
		Status:   1,
	}

	response.Created(c, user)
}

// Login godoc
// @Summary User login
// @Tags User
// @Accept json
// @Produce json
// @Param body body LoginReq true "Login credentials"
// @Success 200 {object} LoginResp
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// TODO: verify password, generate JWT tokens
	resp := &LoginResp{
		AccessToken:  "mock_access_token",
		RefreshToken: "mock_refresh_token",
		User: &model.User{
			Username: "test_user",
			Email:    req.Email,
			Role:     "student",
		},
	}

	response.OK(c, resp)
}

// GetProfile godoc
// @Summary Get current user profile
// @Tags User
// @Security BearerAuth
// @Success 200 {object} model.User
// @Router /api/v1/users/me [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetInt64("user_id")
	_ = userID // TODO: fetch from database
	response.OK(c, gin.H{"user_id": userID})
}

// UpdateProfile godoc
// @Summary Update user profile
// @Tags User
// @Security BearerAuth
// @Param body body UpdateUserReq true "Update data"
// @Success 200
// @Router /api/v1/users/me [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	userID := c.GetInt64("user_id")
	_ = userID // TODO: update in database
	response.OK(c, nil)
}

// ==================== Helper ====================

func getUserID(c *gin.Context) int64 {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}
