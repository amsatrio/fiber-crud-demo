package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("AUTH_SECRET_TOKEN"))

type AuthService interface {
	Login(username string, password string) (*Auth, error)
	Register(username string, password string) error
	RefreshToken(token string) (*Auth, error)
	ForgotPassword(username string) error
	ResetPassword(username string, password string) error
}

type AuthServiceImpl struct {
	repo AuthRepository
}

// ForgotPassword implements [AuthService].
func (a *AuthServiceImpl) ForgotPassword(username string) error {
	panic("unimplemented")
}

// Login implements [AuthService].
func (a *AuthServiceImpl) Login(username string, password string) (*Auth, error) {
	// Verify user credentials via repository
	mUser, err := a.repo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	// Define token expiration times
	accessTokenExpiry := time.Now().Add(time.Hour * 1).Unix()       // 1 hour
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days

	mainClaims := jwt.MapClaims{
		"username": mUser.Email,
		"role":     mUser.RoleId,
		"exp":      accessTokenExpiry,
	}
	mainTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mainClaims)
	mainTokenStr, err := mainTokenObj.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token
	refreshClaims := jwt.MapClaims{
		"username": mUser.Email,
		"role":     mUser.RoleId,
		"exp":      refreshTokenExpiry,
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh_tokenStr, err := refreshTokenObj.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Map values to your Auth DTO
	auth := &Auth{
		MainToken:    &mainTokenStr,
		RefreshToken: &refresh_tokenStr,
		ExpiredIn:    &accessTokenExpiry,
	}

	return auth, nil
}

// RefreshToken implements [AuthService].
func (a *AuthServiceImpl) RefreshToken(token string) (*Auth, error) {
	panic("unimplemented")
}

// Register implements [AuthService].
func (a *AuthServiceImpl) Register(username string, password string) error {
	panic("unimplemented")
}

// ResetPassword implements [AuthService].
func (a *AuthServiceImpl) ResetPassword(username string, password string) error {
	panic("unimplemented")
}

func NewAuthService(repo AuthRepository) AuthService {
	return &AuthServiceImpl{
		repo: repo,
	}
}
