package userBiz

import (
	"backend_autotest/common"
	"backend_autotest/component/generatetoken"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginUserStore interface {
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type loginUserBiz struct {
	store CreateUserStore
}

func NewLoginUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store}
}

func (biz *createUserBiz) LoginUser(ctx context.Context, data *userModel.UserLogin, secret string) (string, error) {

	user, err := biz.store.FindUser(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		if err.Error() != common.RecordNotFound {
			return "", common.ErrDB(err)
		}
		return "", common.ErrInvalidLogin(err)
	}

	if user.Password != data.Password {
		return "", common.ErrInvalidLogin(err)
	}

	payload := generatetoken.Payload{data.UserName}
	provider := generatetoken.NewProviderToken(payload, secret)
	token, err := provider.Encrypt()
	if err != nil {
		return "", common.ErrInvalidLogin(err)
	}

	return token, nil
}
