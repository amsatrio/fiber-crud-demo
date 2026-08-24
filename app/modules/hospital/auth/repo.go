package auth

import (
	"errors"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/modules/hospital/m_user"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsernameAndPassword(username string, password string) (*m_user.MUser, error)
	FindByUsername(username string) (*m_user.MUser, error)
}

type AuthRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

// FindByUsernameAndPassword implements [AuthRepository].
func (a *AuthRepositoryImpl) FindByUsernameAndPassword(username string, password string) (*m_user.MUser, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	var entity m_user.MUser

	result := a.db.Where("email = ? AND password = ?", username, password).First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, result.Error
	}

	return &entity, nil
}

// FindByUsername implements [AuthRepository].
func (a *AuthRepositoryImpl) FindByUsername(username string) (*m_user.MUser, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	var entity m_user.MUser

	result := a.db.Where("email = ?", username).First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username")
		}
		return nil, result.Error
	}

	return &entity, nil
}

func NewMAdminRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}
