package userStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userModel"
	"context"
)

func (db *mongoStore) CreateUser(ctx context.Context, data *userModel.User) error {
	collection := db.db.Database("AutomationTest").Collection("User")

	if _, err := collection.InsertOne(ctx, data); err != nil {
		component.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return common.ErrDB(err)

	}
	return nil
}
