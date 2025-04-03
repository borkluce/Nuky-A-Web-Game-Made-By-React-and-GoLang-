package repo

import (
	"errors"

	"gorm.io/gorm"

	"services/internal/auth/model"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (ap *AuthRepo) CreateUser(user *model.User) error {
	if err := ap.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ap *AuthRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ap.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ap *AuthRepo) GetUserByID(userID int) (*model.User, error) {
	var user model.User
	if err := ap.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ap *AuthRepo) UpdateLastMoveDate(userID int, lastMoveDate string) error {
	if err := ap.db.Model(&model.User{}).Where("id = ?", userID).Update("last_move_date", lastMoveDate).Error; err != nil {
		return err
	}
	return nil
}
