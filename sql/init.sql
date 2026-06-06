-- Smart Online Education Platform - Database Schema
-- 智能在线教育平台 - 数据库建表脚本

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users
CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    username        VARCHAR(50) NOT NULL UNIQUE,
    email           VARCHAR(255) NOT NULL UNIQUE,
    phone           VARCHAR(20),
    password_hash   VARCHAR(255) NOT NULL,
    nickname        VARCHAR(50),
    avatar_url      VARCHAR(512),
    role            VARCHAR(20) NOT NULL DEFAULT 'student' CHECK (role IN ('student','teacher','admin')),
    status          SMALLINT NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_users_role ON users(role);

-- Categories
CREATE TABLE categories (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    parent_id   BIGINT REFERENCES categories(id),
    sort_order  INT NOT NULL DEFAULT 0,
    icon        VARCHAR(512)
);

-- Courses
CREATE TABLE courses (
    id              BIGSERIAL PRIMARY KEY,
    teacher_id      BIGINT NOT NULL REFERENCES users(id),
    title           VARCHAR(256) NOT NULL,
    description     TEXT,
    cover_image     VARCHAR(512),
    category_id     BIGINT REFERENCES categories(id),
    price           DECIMAL(10,2) NOT NULL DEFAULT 0,
    level           VARCHAR(20) NOT NULL DEFAULT 'beginner' CHECK (level IN ('beginner','intermediate','advanced')),
    status          SMALLINT NOT NULL DEFAULT 0,
    student_count   INT NOT NULL DEFAULT 0,
    rating          DECIMAL(3,2) DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_courses_teacher ON courses(teacher_id);
CREATE INDEX idx_courses_category ON courses(category_id);
CREATE INDEX idx_courses_status ON courses(status);

-- Chapters
CREATE TABLE chapters (
    id          BIGSERIAL PRIMARY KEY,
    course_id   BIGINT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title       VARCHAR(256) NOT NULL,
    sort_order  INT NOT NULL DEFAULT 0,
    type        VARCHAR(20) NOT NULL DEFAULT 'video' CHECK (type IN ('video','document','live')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_chapters_course ON chapters(course_id);

-- Videos
CREATE TABLE videos (
    id          BIGSERIAL PRIMARY KEY,
    chapter_id  BIGINT NOT NULL REFERENCES chapters(id) ON DELETE CASCADE,
    title       VARCHAR(256) NOT NULL,
    duration    INT NOT NULL DEFAULT 0,
    url         VARCHAR(512),
    hls_url     VARCHAR(512),
    status      SMALLINT NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Exams
CREATE TABLE exams (
    id          BIGSERIAL PRIMARY KEY,
    course_id   BIGINT NOT NULL REFERENCES courses(id),
    title       VARCHAR(256) NOT NULL,
    description TEXT,
    duration    INT NOT NULL DEFAULT 60,
    pass_score  DECIMAL(5,2) NOT NULL DEFAULT 60.00,
    status      SMALLINT NOT NULL DEFAULT 1,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Questions
CREATE TABLE questions (
    id          BIGSERIAL PRIMARY KEY,
    exam_id     BIGINT NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
    type        VARCHAR(20) NOT NULL DEFAULT 'single_choice',
    content     TEXT NOT NULL,
    options     JSONB,
    answer      VARCHAR(500) NOT NULL,
    score       DECIMAL(5,2) NOT NULL DEFAULT 1.00,
    sort_order  INT NOT NULL DEFAULT 0
);

-- Orders
CREATE TABLE orders (
    id              BIGSERIAL PRIMARY KEY,
    order_no        VARCHAR(64) NOT NULL UNIQUE,
    user_id         BIGINT NOT NULL REFERENCES users(id),
    course_id       BIGINT NOT NULL REFERENCES courses(id),
    amount          DECIMAL(10,2) NOT NULL,
    discount_amount DECIMAL(10,2) NOT NULL DEFAULT 0,
    pay_amount      DECIMAL(10,2) NOT NULL,
    pay_channel     VARCHAR(32),
    status          SMALLINT NOT NULL DEFAULT 0,
    pay_time        TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_orders_user ON orders(user_id);
CREATE INDEX idx_orders_course ON orders(course_id);

-- Learning progress
CREATE TABLE learning_progress (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id),
    course_id   BIGINT NOT NULL REFERENCES courses(id),
    chapter_id  BIGINT NOT NULL REFERENCES chapters(id),
    progress    SMALLINT NOT NULL DEFAULT 0,
    position    INT NOT NULL DEFAULT 0,
    completed   BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, chapter_id)
);

-- Certificates
CREATE TABLE certificates (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id),
    course_id   BIGINT NOT NULL REFERENCES courses(id),
    cert_no     VARCHAR(64) NOT NULL UNIQUE,
    issued_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, course_id)
);

-- Coupons
CREATE TABLE coupons (
    id              BIGSERIAL PRIMARY KEY,
    code            VARCHAR(32) NOT NULL UNIQUE,
    discount_value  DECIMAL(10,2) NOT NULL,
    min_amount      DECIMAL(10,2) NOT NULL DEFAULT 0,
    total_count     INT NOT NULL,
    remain_count    INT NOT NULL,
    start_time      TIMESTAMPTZ NOT NULL,
    end_time        TIMESTAMPTZ NOT NULL,
    status          SMALLINT NOT NULL DEFAULT 1
);

-- Insert default admin
INSERT INTO users (username, email, password_hash, nickname, role, status)
VALUES ('admin', 'admin@education.com', '$2a$10$placeholder_hash', 'System Admin', 'admin', 1);

-- Insert sample categories
INSERT INTO categories (name, sort_order) VALUES
    ('Programming', 1),
    ('Design', 2),
    ('Business', 3),
    ('Language', 4),
    ('Science', 5);
