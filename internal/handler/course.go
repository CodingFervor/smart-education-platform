package handler

import (
	"strconv"

	"github.com/CodingFervor/smart-education-platform/internal/model"
	"github.com/CodingFervor/smart-education-platform/pkg/response"
	"github.com/gin-gonic/gin"
)

// ==================== DTOs ====================

type CreateCourseReq struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	CoverImage  string  `json:"cover_image"`
	CategoryID  int64   `json:"category_id" binding:"required"`
	Price       float64 `json:"price"`
	Level       string  `json:"level" binding:"required,oneof=beginner intermediate advanced"`
}

type UpdateCourseReq struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CoverImage  string  `json:"cover_image"`
	CategoryID  int64   `json:"category_id"`
	Price       float64 `json:"price"`
	Level       string  `json:"level"`
	Status      *int    `json:"status"`
}

type CreateChapterReq struct {
	Title     string `json:"title" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=video document live"`
	SortOrder int    `json:"sort_order"`
}

type ListCoursesReq struct {
	CategoryID *int64  `form:"category_id"`
	Level      string  `form:"level"`
	Keyword    string  `form:"keyword"`
	Page       int     `form:"page,default=1"`
	PageSize   int     `form:"page_size,default=20"`
}

// ==================== Handlers ====================

type CourseHandler struct{}

func NewCourseHandler() *CourseHandler {
	return &CourseHandler{}
}

// CreateCourse godoc
// @Summary Create a new course (teacher only)
// @Tags Course
// @Security BearerAuth
// @Accept json
// @Param body body CreateCourseReq true "Course data"
// @Success 201 {object} model.Course
// @Router /api/v1/courses [post]
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req CreateCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	teacherID := c.GetInt64("user_id")
	course := &model.Course{
		TeacherID:  teacherID,
		Title:      req.Title,
		Description: req.Description,
		CoverImage: req.CoverImage,
		CategoryID: req.CategoryID,
		Price:      req.Price,
		Level:      req.Level,
		Status:     0, // draft
	}

	// TODO: save to database
	response.Created(c, course)
}

// GetCourse godoc
// @Summary Get course detail
// @Tags Course
// @Param id path int true "Course ID"
// @Success 200 {object} model.Course
// @Router /api/v1/courses/{id} [get]
func (h *CourseHandler) GetCourse(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_ = id // TODO: fetch from database
	response.OK(c, gin.H{"id": id, "title": "Mock Course"})
}

// ListCourses godoc
// @Summary List courses with pagination and filters
// @Tags Course
// @Param category_id query int false "Category ID"
// @Param level query string false "Course level"
// @Param keyword query string false "Search keyword"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} response.R
// @Router /api/v1/courses [get]
func (h *CourseHandler) ListCourses(c *gin.Context) {
	var req ListCoursesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// TODO: query from database with pagination
	response.OK(c, gin.H{
		"records": []interface{}{},
		"total":   0,
		"page":    req.Page,
		"size":    req.PageSize,
	})
}

// UpdateCourse godoc
// @Summary Update a course
// @Tags Course
// @Security BearerAuth
// @Param id path int true "Course ID"
// @Param body body UpdateCourseReq true "Update data"
// @Success 200
// @Router /api/v1/courses/{id} [put]
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	var req UpdateCourseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// TODO: update in database
	response.OK(c, nil)
}

// DeleteCourse godoc
// @Summary Delete a course
// @Tags Course
// @Security BearerAuth
// @Param id path int true "Course ID"
// @Success 200
// @Router /api/v1/courses/{id} [delete]
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	// TODO: soft delete
	response.OK(c, nil)
}

// AddChapter godoc
// @Summary Add chapter to course
// @Tags Course
// @Security BearerAuth
// @Param id path int true "Course ID"
// @Param body body CreateChapterReq true "Chapter data"
// @Success 201 {object} model.Chapter
// @Router /api/v1/courses/{id}/chapters [post]
func (h *CourseHandler) AddChapter(c *gin.Context) {
	courseID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req CreateChapterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	chapter := &model.Chapter{
		CourseID:  courseID,
		Title:     req.Title,
		Type:      req.Type,
		SortOrder: req.SortOrder,
	}

	// TODO: save to database
	response.Created(c, chapter)
}
