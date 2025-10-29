# Google Authorization API Documentation

This document provides a comprehensive guide to implementing the Google authorization API endpoints for the egraduate system using Golang and the Gin framework.

## Overview

The Google authorization system allows students to authenticate using their @rumail.ru.ac.th Google accounts. The system handles:
- Student authentication via Google OAuth tokens
- Automatic student registration
- JWT token generation for session management

## API Endpoints

### 1. Google Authorization (Login)
**Endpoint:** `POST /google/authorization`

Authenticates a student using their Google OAuth token and returns JWT tokens for session management.

**Request Body:**
```json
{
  "std_code": "6427952036",
  "refresh_token": "google_oauth_token_here"
}
```

**Response (Success):**
```json
{
  "accessToken": "jwt_access_token",
  "refreshToken": "jwt_refresh_token",
  "isAuth": true
}
```

**Response (Error):**
```json
{
  "error": "Student not found",
  "message": "กรุณาลงทะเบียนก่อนใช้งาน"
}
```

### 2. Student Registration
**Endpoint:** `POST /google/register`

Registers a new student in the system using their Google account information.

**Request Body:**
```json
{
  "std_code": "6427952036",
  "google_token": "google_oauth_token",
  "email": "6427952036@rumail.ru.ac.th",
  "name": "Student Name",
  "picture": "https://profile-picture-url"
}
```

**Response (Success):**
```json
{
  "message": "Registration successful",
  "std_code": "6427952036"
}
```

### 3. Verify Student
**Endpoint:** `GET /google/verify/:studentId`

Checks if a student exists in the system.

**Response (Success):**
```json
{
  "exists": true,
  "std_code": "6427952036"
}
```

**Response (Not Found):**
```json
{
  "exists": false
}
```

## Implementation Guide

### Project Structure

```
project/
├── config/
│   └── database.go
├── models/
│   ├── student.go
│   └── token.go
├── repositories/
│   └── student_repository.go
├── services/
│   ├── google_service.go
│   └── jwt_service.go
├── handlers/
│   └── google_handler.go
├── middleware/
│   └── auth_middleware.go
├── main.go
└── .env
```

### 1. Database Configuration

**File:** `config/database.go`

```go
package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Database connected successfully")
}
```

### 2. Models

**File:** `models/student.go`

```go
package models

import "time"

type Student struct {
    ID          int       `json:"id"`
    StdCode     string    `json:"std_code"`
    Email       string    `json:"email"`
    Name        string    `json:"name"`
    Picture     string    `json:"picture"`
    GoogleToken string    `json:"-"`
    Role        string    `json:"role"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type GoogleAuthRequest struct {
    StdCode      string `json:"std_code" binding:"required"`
    RefreshToken string `json:"refresh_token" binding:"required"`
}

type GoogleRegisterRequest struct {
    StdCode     string `json:"std_code" binding:"required"`
    GoogleToken string `json:"google_token" binding:"required"`
    Email       string `json:"email" binding:"required,email"`
    Name        string `json:"name" binding:"required"`
    Picture     string `json:"picture"`
}
```

**File:** `models/token.go`

```go
package models

type TokenResponse struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
    IsAuth       bool   `json:"isAuth"`
}

type JWTClaims struct {
    StdCode string `json:"std_code"`
    Email   string `json:"email"`
    Role    string `json:"role"`
    Officer string `json:"officer,omitempty"`
}
```

### 3. Repository Layer

**File:** `repositories/student_repository.go`

```go
package repositories

import (
    "database/sql"
    "errors"
    "project/config"
    "project/models"
    "time"
)

type StudentRepository struct {
    db *sql.DB
}

func NewStudentRepository() *StudentRepository {
    return &StudentRepository{db: config.DB}
}

// FindByStdCode retrieves a student by their student code
func (r *StudentRepository) FindByStdCode(stdCode string) (*models.Student, error) {
    query := `
        SELECT id, std_code, email, name, picture, google_token, role, created_at, updated_at
        FROM students
        WHERE std_code = ?
    `

    student := &models.Student{}
    err := r.db.QueryRow(query, stdCode).Scan(
        &student.ID,
        &student.StdCode,
        &student.Email,
        &student.Name,
        &student.Picture,
        &student.GoogleToken,
        &student.Role,
        &student.CreatedAt,
        &student.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return nil, errors.New("student not found")
    }
    if err != nil {
        return nil, err
    }

    return student, nil
}

// Create creates a new student record
func (r *StudentRepository) Create(student *models.Student) error {
    query := `
        INSERT INTO students (std_code, email, name, picture, google_token, role, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `

    now := time.Now()
    result, err := r.db.Exec(
        query,
        student.StdCode,
        student.Email,
        student.Name,
        student.Picture,
        student.GoogleToken,
        student.Role,
        now,
        now,
    )

    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }

    student.ID = int(id)
    student.CreatedAt = now
    student.UpdatedAt = now

    return nil
}

// UpdateGoogleToken updates the student's Google token
func (r *StudentRepository) UpdateGoogleToken(stdCode, token string) error {
    query := `
        UPDATE students
        SET google_token = ?, updated_at = ?
        WHERE std_code = ?
    `

    _, err := r.db.Exec(query, token, time.Now(), stdCode)
    return err
}

// Exists checks if a student exists
func (r *StudentRepository) Exists(stdCode string) (bool, error) {
    query := `SELECT COUNT(*) FROM students WHERE std_code = ?`

    var count int
    err := r.db.QueryRow(query, stdCode).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}
```

### 4. Service Layer

**File:** `services/google_service.go`

```go
package services

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

type GoogleService struct {
    clientID string
}

type GoogleTokenInfo struct {
    Email         string `json:"email"`
    EmailVerified bool   `json:"email_verified"`
    Name          string `json:"name"`
    Picture       string `json:"picture"`
    Sub           string `json:"sub"`
}

func NewGoogleService() *GoogleService {
    return &GoogleService{
        clientID: os.Getenv("GOOGLE_CLIENT_ID"),
    }
}

// VerifyToken verifies the Google OAuth token
func (s *GoogleService) VerifyToken(ctx context.Context, token string) (*GoogleTokenInfo, error) {
    url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", token)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to verify token: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return nil, fmt.Errorf("token verification failed: %s", string(body))
    }

    var tokenInfo GoogleTokenInfo
    if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
        return nil, fmt.Errorf("failed to decode token info: %w", err)
    }

    // Verify email domain
    if !strings.HasSuffix(tokenInfo.Email, "@rumail.ru.ac.th") {
        return nil, errors.New("invalid email domain")
    }

    // Verify email is verified
    if !tokenInfo.EmailVerified {
        return nil, errors.New("email not verified")
    }

    return &tokenInfo, nil
}

// ExtractStudentID extracts student ID from email
func (s *GoogleService) ExtractStudentID(email string) string {
    parts := strings.Split(email, "@")
    if len(parts) > 0 {
        return parts[0]
    }
    return ""
}
```

**File:** `services/jwt_service.go`

```go
package services

import (
    "errors"
    "os"
    "project/models"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
    secretKey string
}

func NewJWTService() *JWTService {
    return &JWTService{
        secretKey: os.Getenv("JWT_SECRET_KEY"),
    }
}

// GenerateTokens generates access and refresh tokens
func (s *JWTService) GenerateTokens(claims models.JWTClaims) (*models.TokenResponse, error) {
    // Generate access token (expires in 24 hours)
    accessToken, err := s.generateToken(claims, 24*time.Hour)
    if err != nil {
        return nil, err
    }

    // Generate refresh token (expires in 7 days)
    refreshToken, err := s.generateToken(claims, 7*24*time.Hour)
    if err != nil {
        return nil, err
    }

    return &models.TokenResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        IsAuth:       true,
    }, nil
}

func (s *JWTService) generateToken(claims models.JWTClaims, expiration time.Duration) (string, error) {
    now := time.Now()

    jwtClaims := jwt.MapClaims{
        "std_code": claims.StdCode,
        "email":    claims.Email,
        "role":     claims.Role,
        "iat":      now.Unix(),
        "exp":      now.Add(expiration).Unix(),
    }

    if claims.Officer != "" {
        jwtClaims["officer"] = claims.Officer
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
    return token.SignedString([]byte(s.secretKey))
}

// VerifyToken verifies and parses a JWT token
func (s *JWTService) VerifyToken(tokenString string) (*models.JWTClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return []byte(s.secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

    jwtClaims := &models.JWTClaims{
        StdCode: claims["std_code"].(string),
        Email:   claims["email"].(string),
        Role:    claims["role"].(string),
    }

    if officer, ok := claims["officer"].(string); ok {
        jwtClaims.Officer = officer
    }

    return jwtClaims, nil
}
```

### 5. Handler Layer

**File:** `handlers/google_handler.go`

```go
package handlers

import (
    "log"
    "net/http"
    "project/models"
    "project/repositories"
    "project/services"

    "github.com/gin-gonic/gin"
)

type GoogleHandler struct {
    studentRepo  *repositories.StudentRepository
    googleSvc    *services.GoogleService
    jwtSvc       *services.JWTService
}

func NewGoogleHandler() *GoogleHandler {
    return &GoogleHandler{
        studentRepo:  repositories.NewStudentRepository(),
        googleSvc:    services.NewGoogleService(),
        jwtSvc:       services.NewJWTService(),
    }
}

// Authorization handles Google OAuth login
func (h *GoogleHandler) Authorization(c *gin.Context) {
    var req models.GoogleAuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "message": "กรุณาระบุข้อมูลให้ครบถ้วน",
        })
        return
    }

    // Verify Google token
    tokenInfo, err := h.googleSvc.VerifyToken(c.Request.Context(), req.RefreshToken)
    if err != nil {
        log.Printf("Token verification failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":   "Token verification failed",
            "message": "ไม่สามารถยืนยันตัวตนได้",
        })
        return
    }

    // Verify student ID matches email
    extractedID := h.googleSvc.ExtractStudentID(tokenInfo.Email)
    if extractedID != req.StdCode {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Student ID mismatch",
            "message": "รหัสนักศึกษาไม่ตรงกับอีเมล",
        })
        return
    }

    // Find student in database
    student, err := h.studentRepo.FindByStdCode(req.StdCode)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "Student not found",
            "message": "ไม่พบข้อมูลนักศึกษา กรุณาลงทะเบียนก่อนใช้งาน",
        })
        return
    }

    // Update Google token
    if err := h.studentRepo.UpdateGoogleToken(req.StdCode, req.RefreshToken); err != nil {
        log.Printf("Failed to update token: %v", err)
    }

    // Generate JWT tokens
    jwtClaims := models.JWTClaims{
        StdCode: student.StdCode,
        Email:   student.Email,
        Role:    student.Role,
    }

    tokens, err := h.jwtSvc.GenerateTokens(jwtClaims)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Token generation failed",
            "message": "เกิดข้อผิดพลาดในการสร้าง token",
        })
        return
    }

    c.JSON(http.StatusOK, tokens)
}

// Register handles student registration
func (h *GoogleHandler) Register(c *gin.Context) {
    var req models.GoogleRegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "message": "กรุณาระบุข้อมูลให้ครบถ้วน",
        })
        return
    }

    // Verify Google token
    tokenInfo, err := h.googleSvc.VerifyToken(c.Request.Context(), req.GoogleToken)
    if err != nil {
        log.Printf("Token verification failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":   "Token verification failed",
            "message": "ไม่สามารถยืนยันตัวตนได้",
        })
        return
    }

    // Verify email matches
    if tokenInfo.Email != req.Email {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Email mismatch",
            "message": "อีเมลไม่ตรงกัน",
        })
        return
    }

    // Check if student already exists
    exists, err := h.studentRepo.Exists(req.StdCode)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Database error",
            "message": "เกิดข้อผิดพลาดในการตรวจสอบข้อมูล",
        })
        return
    }

    if exists {
        c.JSON(http.StatusConflict, gin.H{
            "error":   "Student already exists",
            "message": "นักศึกษาถูกลงทะเบียนแล้ว",
        })
        return
    }

    // Determine role based on student ID pattern
    role := h.determineRole(req.StdCode)

    // Create new student
    student := &models.Student{
        StdCode:     req.StdCode,
        Email:       req.Email,
        Name:        req.Name,
        Picture:     req.Picture,
        GoogleToken: req.GoogleToken,
        Role:        role,
    }

    if err := h.studentRepo.Create(student); err != nil {
        log.Printf("Failed to create student: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Registration failed",
            "message": "ไม่สามารถลงทะเบียนได้",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message":  "Registration successful",
        "std_code": student.StdCode,
    })
}

// VerifyStudent checks if a student exists
func (h *GoogleHandler) VerifyStudent(c *gin.Context) {
    studentID := c.Param("studentId")

    exists, err := h.studentRepo.Exists(studentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Database error",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "exists":   exists,
        "std_code": studentID,
    })
}

// determineRole determines student role based on student ID
func (h *GoogleHandler) determineRole(stdCode string) string {
    // This is a simple implementation
    // You may need to adjust based on your university's student ID format
    if len(stdCode) >= 2 {
        // Example: Master students might have specific patterns
        // Adjust this logic based on your requirements
        return "Master" // Default role
    }
    return "Master"
}
```

### 6. Main Application

**File:** `main.go`

```go
package main

import (
    "log"
    "os"
    "project/config"
    "project/handlers"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    config.InitDatabase()
    defer config.DB.Close()

    // Setup Gin router
    router := gin.Default()

    // CORS configuration
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Initialize handlers
    googleHandler := handlers.NewGoogleHandler()

    // Google authentication routes
    googleRoutes := router.Group("/google")
    {
        googleRoutes.POST("/authorization", googleHandler.Authorization)
        googleRoutes.POST("/register", googleHandler.Register)
        googleRoutes.GET("/verify/:studentId", googleHandler.VerifyStudent)
    }

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
```

### 7. Environment Configuration

**File:** `.env`

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=egraduate

# JWT Configuration
JWT_SECRET_KEY=your-super-secret-jwt-key-change-this-in-production

# Google OAuth Configuration
GOOGLE_CLIENT_ID=505837308278-q7h4ihqtu9quajaotvtejhosb2pepmmt.apps.googleusercontent.com

# Server Configuration
PORT=8080
FRONTEND_URL=http://localhost:4200

# Environment
ENV=development
```

### 8. Database Schema

**File:** `database/schema.sql`

```sql
CREATE TABLE IF NOT EXISTS students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    std_code VARCHAR(10) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    picture TEXT,
    google_token TEXT,
    role ENUM('Master', 'Doctor', 'admin') DEFAULT 'Master',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_std_code (std_code),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## Required Dependencies

**File:** `go.mod`

```go
module project

go 1.21

require (
    github.com/gin-contrib/cors v1.5.0
    github.com/gin-gonic/gin v1.9.1
    github.com/go-sql-driver/mysql v1.7.1
    github.com/golang-jwt/jwt/v4 v4.5.0
    github.com/joho/godotenv v1.5.1
)
```

## Installation Steps

1. **Initialize Go Module:**
```bash
go mod init project
go mod tidy
```

2. **Install Dependencies:**
```bash
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/go-sql-driver/mysql
go get github.com/golang-jwt/jwt/v4
go get github.com/joho/godotenv
```

3. **Create Database:**
```bash
mysql -u root -p < database/schema.sql
```

4. **Configure Environment:**
- Copy `.env.example` to `.env`
- Update database credentials
- Set JWT secret key
- Configure Google Client ID

5. **Run Application:**
```bash
go run main.go
```

## Testing with cURL

### Test Authorization
```bash
curl -X POST http://localhost:8080/google/authorization \
  -H "Content-Type: application/json" \
  -d '{
    "std_code": "6427952036",
    "refresh_token": "your_google_oauth_token"
  }'
```

### Test Registration
```bash
curl -X POST http://localhost:8080/google/register \
  -H "Content-Type: application/json" \
  -d '{
    "std_code": "6427952036",
    "google_token": "your_google_oauth_token",
    "email": "6427952036@rumail.ru.ac.th",
    "name": "Test Student",
    "picture": "https://profile-pic-url"
  }'
```

### Test Verify
```bash
curl http://localhost:8080/google/verify/6427952036
```

## Security Considerations

1. **Token Security:**
   - Always verify Google tokens on the server side
   - Use strong JWT secret keys (minimum 32 characters)
   - Set appropriate token expiration times
   - Never expose secret keys in client-side code

2. **Email Domain Validation:**
   - Always verify @rumail.ru.ac.th domain
   - Check email verification status from Google

3. **HTTPS:**
   - Always use HTTPS in production
   - Configure TLS certificates properly

4. **Database Security:**
   - Use prepared statements (already implemented)
   - Encrypt sensitive data at rest
   - Regular backups

5. **Rate Limiting:**
   - Implement rate limiting for authentication endpoints
   - Use middleware like `gin-rate-limit`

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Authentication failed
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource already exists
- `500 Internal Server Error` - Server error

## Deployment Checklist

- [ ] Update `.env` with production values
- [ ] Set strong JWT_SECRET_KEY
- [ ] Configure CORS for production domain
- [ ] Enable HTTPS/TLS
- [ ] Set up database backups
- [ ] Configure logging
- [ ] Set up monitoring
- [ ] Implement rate limiting
- [ ] Review security settings
- [ ] Test all endpoints thoroughly

## Additional Resources

- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)
- [Google OAuth 2.0 Documentation](https://developers.google.com/identity/protocols/oauth2)
- [JWT Best Practices](https://tools.ietf.org/html/rfc8725)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)

## Support

For issues or questions, please contact the development team or refer to the project documentation.
