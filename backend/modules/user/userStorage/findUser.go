package userStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error) {
	collection := db.db.Database("AutomationTest").Collection("User")

	var data bson.M

	if err := collection.FindOne(ctx, conditions).Decode(&data); err != nil {
		if err.Error() == common.RecordNotFound {
			return nil, err
		}
		return nil, common.ErrDB(err)
	}

	var result userModel.User
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &result)
	return &result, nil
}
