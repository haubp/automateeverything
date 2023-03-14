package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *mongoStore) PostResult(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error) {
	collection := db.db.Database("AutomationTest").Collection("Node_Result")

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		component.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return nil, common.ErrDB(err)
	}

	return res, nil
}
