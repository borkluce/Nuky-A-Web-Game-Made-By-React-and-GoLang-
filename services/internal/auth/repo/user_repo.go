package repo

import (
	"errors"

	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (ap *AuthRepo) CreateUser(user *models.User) error {
	if err := ap.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ap *AuthRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ap.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
