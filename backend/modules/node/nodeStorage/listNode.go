package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) ListNode(ctx context.Context, conditions map[string]interface{}) ([]nodeModel.Node, error) {
	collection := db.db.Database("AutomationTest").Collection("Node")

	cursor, err := collection.Find(ctx, conditions)
	if err != nil {
		return nil, common.ErrDB(err)
	}

	var convert []bson.D
	if err = cursor.All(context.TODO(), &convert); err != nil {
		common.ErrInternal(err)
	}

	var results []nodeModel.Node
	for _, result := range convert {
		var node nodeModel.Node
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &node)
		results = append(results, node)
	}

	return results, nil
}
