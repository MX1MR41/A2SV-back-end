package Tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager/Infrastructure"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

// Define the test suites
type PasswordServiceTestSuite struct {
	suite.Suite
}

type JWTServiceTestSuite struct {
	suite.Suite
}

type AuthMiddlewareTestSuite struct {
	suite.Suite
	router *gin.Engine
}

// Setup the test suites
func (suite *AuthMiddlewareTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()

	// Register routes once in SetupSuite
	suite.router.POST("/login", Infrastructure.Login("test_task_manager"))
	suite.router.GET("/logged", Infrastructure.Logged)
	suite.router.GET("/admin", Infrastructure.Admin)
}

// PasswordServiceTestSuite tests
func (suite *PasswordServiceTestSuite) TestComparePasswords_NoMatch() {
	plainPassword := "password1"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("differentPassword"), bcrypt.DefaultCost)

	err := Infrastructure.ComparePasswords(string(hashedPassword), plainPassword)
	suite.Error(err)
	assert.Equal(suite.T(), "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())
}

func (suite *PasswordServiceTestSuite) TestComparePasswords_Match() {
	plainPassword := "password1"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	err := Infrastructure.ComparePasswords(string(hashedPassword), plainPassword)
	suite.NoError(err)
}

func (suite *PasswordServiceTestSuite) TestComparePasswords_EmptyPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)

	err := Infrastructure.ComparePasswords(string(hashedPassword), "")
	suite.Error(err)
	assert.Equal(suite.T(), "password cannot be empty", err.Error())
}

// JWTServiceTestSuite tests
func (suite *JWTServiceTestSuite) TestGenerateToken_Success() {
	token, err := Infrastructure.GenerateToken("testuser", "user")
	suite.NoError(err)
	suite.NotEmpty(token)
}

func (suite *JWTServiceTestSuite) TestValidateToken_Success() {
	tokenString, _ := Infrastructure.GenerateToken("testuser", "user")
	token, err := Infrastructure.ValidateToken(tokenString)
	suite.NoError(err)
	suite.True(token.Valid)
}

func (suite *JWTServiceTestSuite) TestValidateToken_InvalidToken() {
	token, err := Infrastructure.ValidateToken("invalidtoken")
	suite.Error(err)
	suite.Nil(token)
}

// AuthMiddlewareTestSuite tests
func (suite *AuthMiddlewareTestSuite) TestLogin_InvalidPayload() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *AuthMiddlewareTestSuite) TestLogin_WrongPassword() {
	w := httptest.NewRecorder()
	body := `{"username":"testuser","password":"wrongpassword"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *AuthMiddlewareTestSuite) TestLogged_NoAuthHeader() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logged", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *AuthMiddlewareTestSuite) TestLogged_InvalidAuthHeader() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logged", nil)
	req.Header.Set("Authorization", "InvalidHeader")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *AuthMiddlewareTestSuite) TestLogged_InvalidToken() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logged", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *AuthMiddlewareTestSuite) TestAdmin_Unauthorized() {
	w := httptest.NewRecorder()
	token, _ := Infrastructure.GenerateToken("testuser", "user")
	req, _ := http.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusForbidden, w.Code)
}

// Run each suite independently
func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}

func TestJWTServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
