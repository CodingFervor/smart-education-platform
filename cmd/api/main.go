package main

import (
	"log"

	"github.com/CodingFervor/smart-education-platform/internal/config"
	"github.com/CodingFervor/smart-education-platform/internal/handler"
	"github.com/CodingFervor/smart-education-platform/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup Gin
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Global middleware
	r.Use(middleware.CORS())

	// Handlers
	userH := handler.NewUserHandler()
	courseH := handler.NewCourseHandler()
	examH := handler.NewExamHandler()
	orderH := handler.NewOrderHandler()

	// ==================== Public Routes ====================
	pub := r.Group("/api/v1")
	{
		auth := pub.Group("/auth")
		{
			auth.POST("/register", userH.Register)
			auth.POST("/login", userH.Login)
		}

		// Public course listing
		pub.GET("/courses", courseH.ListCourses)
		pub.GET("/courses/:id", courseH.GetCourse)
	}

	// ==================== Authenticated Routes ====================
	authed := r.Group("/api/v1")
	authed.Use(middleware.AuthMiddleware(&cfg.JWT))
	{
		// User
		authed.GET("/users/me", userH.GetProfile)
		authed.PUT("/users/me", userH.UpdateProfile)

		// Course management (teacher)
		teacher := authed.Group("")
		teacher.Use(middleware.RoleMiddleware("teacher"))
		{
			teacher.POST("/courses", courseH.CreateCourse)
			teacher.PUT("/courses/:id", courseH.UpdateCourse)
			teacher.DELETE("/courses/:id", courseH.DeleteCourse)
			teacher.POST("/courses/:id/chapters", courseH.AddChapter)
			teacher.POST("/exams", examH.CreateExam)
		}

		// Exam
		authed.GET("/exams/:id", examH.GetExam)
		authed.POST("/exams/:id/submit", examH.SubmitExam)

		// Order
		authed.POST("/orders", orderH.CreateOrder)
		authed.GET("/orders", orderH.ListMyOrders)
		authed.GET("/orders/:order_no", orderH.GetOrder)

		// Learning progress
		authed.GET("/learning/:course_id/progress", orderH.GetLearningProgress)
		authed.PUT("/learning/:course_id/progress", orderH.UpdateProgress)
	}

	// Payment callback (no auth)
	r.POST("/api/v1/pay/callback/wechat", orderH.PayCallback)
	r.POST("/api/v1/pay/callback/alipay", orderH.PayCallback)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "education-platform"})
	})

	// Start server
	addr := ":8080"
	log.Printf("Education Platform starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
