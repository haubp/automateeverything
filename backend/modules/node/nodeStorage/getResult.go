package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) GetResult(ctx context.Context, conditions interface{}) (*nodeModel.Result, error) {

	var data bson.M

	collection := db.db.Database("AutomationTest").Collection("Node_Result")
	if err := collection.FindOne(ctx, conditions).Decode(&data); err != nil {
		if err.Error() == common.RecordNotFound {
			return nil, err
		}

		component.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return nil, common.ErrDB(err)
	}

	var result nodeModel.Result
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &result)
	return &result, nil
}
