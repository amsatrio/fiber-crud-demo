package auth

import (
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/user"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsernameOrEmail(identifier string) (*user.User, error)
	FindByUsername(username string) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}

type AuthRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

// FindByUsernameOrEmail implements [AuthRepository].
func (a *AuthRepositoryImpl) FindByUsernameOrEmail(identifier string) (*user.User, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	var entity user.User

	result := a.db.Where("username = ? OR email = ?", identifier, identifier).First(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

// FindByUsername implements [AuthRepository].
func (a *AuthRepositoryImpl) FindByUsername(username string) (*user.User, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	var entity user.User

	result := a.db.Where("username = ?", username).First(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

// FindByEmail implements [AuthRepository].
func (a *AuthRepositoryImpl) FindByEmail(email string) (*user.User, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	var entity user.User

	result := a.db.Where("email = ?", email).First(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}
