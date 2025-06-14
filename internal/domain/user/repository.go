package user

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, users *[]User, req filter.FilterRequest, clientID uint, superAdminID uint) (int, error)
	GetUserByID(ctx context.Context, domainUser *User, userID int, clientID uint, superAdminID uint) error
	GetMe(ctx context.Context, domainUser *User, userID int, clientID uint, isAdmin bool) error
	GetCommonUserByID(ctx context.Context, domainUser *User, userID int) error
	GetByEmailPhoneNumber(ctx context.Context, user *User, email string, phoneNumber string) error
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, domainUser *User) error
	IsExist(ctx context.Context, userID uint) (bool, error)
	IsEmailExist(ctx context.Context, email string, userID int) (bool, error)
	IsPhoneNumberExist(ctx context.Context, phoneNumber string, userID int) (bool, error)
	IsTableEmpty(ctx context.Context) (bool, error)
}
