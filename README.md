# Smart Online Education Platform | 智能在线教育平台

[English](#english) | [中文](#中文)

---

## English

### Overview

A full-featured online education platform built with Go, supporting course management, video-on-demand, live streaming, online examinations, learning progress tracking, and certificate issuance.

### Tech Stack

| Component | Technology |
|-----------|-----------|
| Language | Go 1.22 |
| Web Framework | Gin |
| Database | PostgreSQL + Redis |
| Message Queue | NATS |
| Object Storage | MinIO |
| Authentication | JWT |
| API Docs | Swagger (swaggo) |
| Container | Docker + Docker Compose |

### Architecture

```
┌─────────────── Client Layer ───────────────┐
│   Web App  │  Mobile App  │  Admin Panel   │
└───────────────────┬────────────────────────┘
                    │
              ┌─────┴─────┐
              │ API Gateway│ (Gin + Middleware)
              └─────┬─────┘
                    │
    ┌───────┬───────┼───────┬────────┐
    ▼       ▼       ▼       ▼        ▼
 User    Course   Video   Exam    Order
 Service  Service  Service Service Service
    │       │       │       │        │
    └───────┴───────┴───────┴────────┘
                    │
    ┌───────┬───────┼───────┐
    ▼       ▼       ▼       ▼
PostgreSQL Redis  MinIO   NATS
```

### Quick Start

```bash
# Clone
git clone https://github.com/CodingFervor/smart-education-platform.git
cd smart-education-platform

# Start dependencies
docker-compose up -d

# Run the server
go run cmd/api/main.go

# API docs
open http://localhost:8080/swagger/index.html
```

### Project Structure

```
smart-education-platform/
├── cmd/api/              # Application entry point
├── internal/
│   ├── config/           # Configuration
│   ├── handler/          # HTTP handlers (controllers)
│   ├── middleware/        # Auth, CORS, logging
│   ├── model/            # Data models
│   ├── repository/       # Database access layer
│   └── service/          # Business logic
├── pkg/
│   ├── database/         # DB connection
│   ├── jwt/              # JWT utilities
│   ├── logger/           # Logging
│   └── response/         # Unified response
├── configs/              # Config files
├── docs/                 # Architecture & API docs
├── sql/                  # Database migrations
├── scripts/              # Build & deploy scripts
├── docker-compose.yml
├── Makefile
└── go.mod
```

---

## 中文

### 项目简介

基于 Go 构建的全功能在线教育平台，支持课程管理、视频点播、直播授课、在线考试、学习进度追踪和证书颁发。

### 技术栈

| 组件 | 技术 |
|------|------|
| 语言 | Go 1.22 |
| Web 框架 | Gin |
| 数据库 | PostgreSQL + Redis |
| 消息队列 | NATS |
| 对象存储 | MinIO |
| 认证 | JWT |
| API 文档 | Swagger (swaggo) |
| 容器化 | Docker + Docker Compose |

### 快速开始

```bash
# 克隆项目
git clone https://github.com/CodingFervor/smart-education-platform.git
cd smart-education-platform

# 启动依赖
docker-compose up -d

# 运行服务
go run cmd/api/main.go

# API 文档
open http://localhost:8080/swagger/index.html
```

### 核心功能

- **用户管理**：注册/登录、角色权限（学员/讲师/管理员）
- **课程管理**：课程CRUD、章节管理、视频上传
- **视频点播**：HLS转码、进度记录、断点续播
- **在线考试**：题库管理、自动组卷、自动评分
- **订单支付**：课程购买、优惠券、微信/支付宝
- **学习追踪**：学习进度、完课证书
- **直播授课**：白板互动、实时聊天

---

## License

MIT License
