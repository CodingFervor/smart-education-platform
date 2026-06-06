package handler

import (
	"strconv"

	"github.com/CodingFervor/smart-education-platform/pkg/response"
	"github.com/gin-gonic/gin"
)

type CreateExamReq struct {
	CourseID    int64   `json:"course_id" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Duration    int     `json:"duration" binding:"required,min=1"` // minutes
	PassScore   float64 `json:"pass_score" binding:"required"`
}

type SubmitExamReq struct {
	Answers map[string]string `json:"answers" binding:"required"` // question_id -> answer
}

type ExamHandler struct{}

func NewExamHandler() *ExamHandler {
	return &ExamHandler{}
}

// CreateExam godoc
// @Summary Create an exam (teacher only)
// @Tags Exam
// @Security BearerAuth
// @Accept json
// @Param body body CreateExamReq true "Exam data"
// @Success 201
// @Router /api/v1/exams [post]
func (h *ExamHandler) CreateExam(c *gin.Context) {
	var req CreateExamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// TODO: save to database
	response.Created(c, req)
}

// GetExam godoc
// @Summary Get exam detail with questions
// @Tags Exam
// @Param id path int true "Exam ID"
// @Success 200
// @Router /api/v1/exams/{id} [get]
func (h *ExamHandler) GetExam(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	response.OK(c, gin.H{"exam_id": id})
}

// SubmitExam godoc
// @Summary Submit exam answers
// @Tags Exam
// @Security BearerAuth
// @Param id path int true "Exam ID"
// @Param body body SubmitExamReq true "Answers"
// @Success 200
// @Router /api/v1/exams/{id}/submit [post]
func (h *ExamHandler) SubmitExam(c *gin.Context) {
	var req SubmitExamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	// TODO: auto-grade and return score
	response.OK(c, gin.H{"score": 85.5, "passed": true})
}
