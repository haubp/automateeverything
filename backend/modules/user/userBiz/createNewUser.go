package userBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *userModel.User) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *userModel.User) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidPassword(err)
	}

	user, err := biz.store.FindUser(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		if err.Error() != common.RecordNotFound {
			return common.ErrDB(err)
		}

	}

	if user != nil {
		return common.ErrEntityExisted("User", nil)
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
