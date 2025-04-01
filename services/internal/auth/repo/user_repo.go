package repo

import "github.com/gofiber/fiber/v2"

type AuthRepo struct {
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (ap *AuthRepo) CreateUser(ctx *fiber.Ctx, user *model.User) {}

func (ap *AuthRepo) GetUser(ctx *fiber.Ctx, userID int8) (*model.User, error) {}
