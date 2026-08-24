package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_user"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("AUTH_SECRET_TOKEN"))

type AuthService interface {
	Login(username string, password string) (*Auth, error)
	Register(username string, password string) error
	RefreshToken(token string) (*Auth, error)
	ResetPassword(username string, password string, mUserId uint) error
}

type AuthServiceImpl struct {
	repo            AuthRepository
	mUserRepository m_user.MUserRepository
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
		"userId":   mUser.Id,
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
		"userId":   mUser.Id,
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
	// 1. Parse and validate the incoming refresh token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches HS256
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	// 2. Extract claims safely
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("missing or invalid username in token")
	}

	// In jwt.MapClaims, numeric values unmarshal into float64 by default
	role, ok := claims["role"]
	if !ok {
		return nil, errors.New("missing role in token")
	}

	userId, ok := claims["userId"]
	if !ok {
		return nil, errors.New("missing userId in token")
	}

	// 3. Define new expiration times
	accessTokenExpiry := time.Now().Add(time.Hour * 1).Unix()       // 1 hour
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days

	// 4. Generate new Access Token
	mainClaims := jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"role":     role,
		"exp":      accessTokenExpiry,
	}
	mainTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mainClaims)
	mainTokenStr, err := mainTokenObj.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// 5. Generate new Refresh Token (Token Rotation)
	refreshClaims := jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"role":     role,
		"exp":      refreshTokenExpiry,
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshTokenObj.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// 6. Map values to your Auth DTO
	auth := &Auth{
		MainToken:    &mainTokenStr,
		RefreshToken: &refreshTokenStr,
		ExpiredIn:    &accessTokenExpiry,
	}

	return auth, nil
}

// Register implements [AuthService].
func (a *AuthServiceImpl) Register(username string, password string) error {
	panic("unimplemented")
}

// ResetPassword implements [AuthService].
func (a *AuthServiceImpl) ResetPassword(username string, password string, mUserId uint) error {
	mUser, err := a.repo.FindByUsername(username)
	if err != nil {
		return err
	}

	mUser.Password = &password
	mUser.ModifiedBy = &mUserId
	mUser.ModifiedOn = &dto.JSONTime{Time: time.Now()}

	err = a.mUserRepository.Update(mUser)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo AuthRepository, mUserRepository m_user.MUserRepository) AuthService {
	return &AuthServiceImpl{
		repo:            repo,
		mUserRepository: mUserRepository,
	}
}
