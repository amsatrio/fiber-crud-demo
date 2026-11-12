package auth

import (
	"errors"
	"os"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("AUTH_SECRET_TOKEN"))

type AuthService interface {
	Login(identifier string, password string) (*Auth, error)
	Register(payload AuthRegisterRequest) error
	RefreshToken(token string) (*Auth, error)
	ResetPassword(identifier string, newPassword string, mUserId uint) error
}

type AuthServiceImpl struct {
	repo            AuthRepository
	mUserRepository user.UserRepository
}

// Login implements [AuthService].
func (a *AuthServiceImpl) Login(identifier string, password string) (*Auth, error) {
	// Verify user credentials via repository
	mUser, err := a.repo.FindByUsernameOrEmail(identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username/email or password")
		}
		return nil, err
	}

	if mUser.PasswordHash == nil || bcrypt.CompareHashAndPassword([]byte(*mUser.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("invalid username/email or password")
	}

	username := ""
	if mUser.Username != nil {
		username = *mUser.Username
	} else if mUser.Email != nil {
		username = *mUser.Email
	}

	// Define token expiration times
	accessTokenExpiry := time.Now().Add(time.Hour * 1).Unix()       // 1 hour
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days

	mainClaims := jwt.MapClaims{
		"username": username,
		"userId":   mUser.Id,
		"exp":      accessTokenExpiry,
	}
	mainTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mainClaims)
	mainTokenStr, err := mainTokenObj.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token
	refreshClaims := jwt.MapClaims{
		"username": username,
		"userId":   mUser.Id,
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
			return nil, errors.New("unexpected signing method")
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
func (a *AuthServiceImpl) Register(payload AuthRegisterRequest) error {
	username := ""
	if payload.Username != nil {
		username = *payload.Username
	}
	email := ""
	if payload.Email != nil {
		email = *payload.Email
	}
	password := ""
	if payload.Password != nil {
		password = *payload.Password
	}
	displayName := ""
	if payload.DisplayName != nil {
		displayName = *payload.DisplayName
	}
	if displayName == "" {
		displayName = username
	}

	// Ensure username is not already taken
	if _, err := a.repo.FindByUsername(username); err == nil {
		return errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Ensure email is not already taken
	if _, err := a.repo.FindByEmail(email); err == nil {
		return errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashStr := string(hash)

	entity := user.User{
		Username:     &username,
		Email:        &email,
		DisplayName:  &displayName,
		PasswordHash: &hashStr,
		CreatedAt:    dto.JSONTime{Time: time.Now()},
	}

	if err := a.mUserRepository.Create(&entity); err != nil {
		return err
	}

	return nil
}

// ResetPassword implements [AuthService].
func (a *AuthServiceImpl) ResetPassword(identifier string, newPassword string, mUserId uint) error {
	mUser, err := a.repo.FindByUsernameOrEmail(identifier)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashStr := string(hash)

	mUser.PasswordHash = &hashStr

	if err := a.mUserRepository.Update(mUser); err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo AuthRepository, mUserRepository user.UserRepository) AuthService {
	return &AuthServiceImpl{
		repo:            repo,
		mUserRepository: mUserRepository,
	}
}
