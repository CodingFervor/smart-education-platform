package model

import "time"

// User represents a platform user (student, teacher, or admin)
type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Password  string    `json:"-" db:"password_hash"`
	Nickname  string    `json:"nickname" db:"nickname"`
	Avatar    string    `json:"avatar" db:"avatar_url"`
	Role      string    `json:"role" db:"role"` // student, teacher, admin
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Course represents a course
type Course struct {
	ID          int64     `json:"id" db:"id"`
	TeacherID   int64     `json:"teacher_id" db:"teacher_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CoverImage  string    `json:"cover_image" db:"cover_image"`
	CategoryID  int64     `json:"category_id" db:"category_id"`
	Price       float64   `json:"price" db:"price"`
	Level       string    `json:"level" db:"level"` // beginner, intermediate, advanced
	Status      int       `json:"status" db:"status"` // 0=draft, 1=published, 2=offline
	StudentCount int      `json:"student_count" db:"student_count"`
	Rating      float64   `json:"rating" db:"rating"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Chapter represents a course chapter
type Chapter struct {
	ID        int64     `json:"id" db:"id"`
	CourseID  int64     `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	SortOrder int       `json:"sort_order" db:"sort_order"`
	Type      string    `json:"type" db:"type"` // video, document, live
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Video represents an uploaded video
type Video struct {
	ID         int64     `json:"id" db:"id"`
	ChapterID  int64     `json:"chapter_id" db:"chapter_id"`
	Title      string    `json:"title" db:"title"`
	Duration   int       `json:"duration" db:"duration"` // seconds
	URL        string    `json:"url" db:"url"`
	HLSURL     string    `json:"hls_url" db:"hls_url"`
	Status     int       `json:"status" db:"status"` // 0=uploading, 1=transcoding, 2=ready
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Exam represents an examination
type Exam struct {
	ID          int64     `json:"id" db:"id"`
	CourseID    int64     `json:"course_id" db:"course_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Duration    int       `json:"duration" db:"duration"` // minutes
	PassScore   float64   `json:"pass_score" db:"pass_score"`
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Question represents an exam question
type Question struct {
	ID        int64             `json:"id" db:"id"`
	ExamID    int64             `json:"exam_id" db:"exam_id"`
	Type      string            `json:"type" db:"type"` // single_choice, multi_choice, true_false, essay
	Content   string            `json:"content" db:"content"`
	Options   map[string]string `json:"options" db:"options"`
	Answer    string            `json:"-" db:"answer"`
	Score     float64           `json:"score" db:"score"`
	SortOrder int               `json:"sort_order" db:"sort_order"`
}

// Order represents a purchase order
type Order struct {
	ID              int64     `json:"id" db:"id"`
	OrderNo         string    `json:"order_no" db:"order_no"`
	UserID          int64     `json:"user_id" db:"user_id"`
	CourseID        int64     `json:"course_id" db:"course_id"`
	Amount          float64   `json:"amount" db:"amount"`
	DiscountAmount  float64   `json:"discount_amount" db:"discount_amount"`
	PayAmount       float64   `json:"pay_amount" db:"pay_amount"`
	PayChannel      string    `json:"pay_channel" db:"pay_channel"`
	Status          int       `json:"status" db:"status"` // 0=unpaid, 1=paid, 2=cancelled, 3=refunded
	PayTime         *time.Time `json:"pay_time" db:"pay_time"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

// LearningProgress tracks student progress
type LearningProgress struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	CourseID   int64     `json:"course_id" db:"course_id"`
	ChapterID  int64     `json:"chapter_id" db:"chapter_id"`
	Progress   int       `json:"progress" db:"progress"` // 0-100
	Position   int       `json:"position" db:"position"` // video position in seconds
	Completed  bool      `json:"completed" db:"completed"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Certificate represents a completion certificate
type Certificate struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	CourseID  int64     `json:"course_id" db:"course_id"`
	CertNo    string    `json:"cert_no" db:"cert_no"`
	IssuedAt  time.Time `json:"issued_at" db:"issued_at"`
}

// Category represents a course category
type Category struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ParentID *int64 `json:"parent_id" db:"parent_id"`
	SortOrder int   `json:"sort_order" db:"sort_order"`
}
