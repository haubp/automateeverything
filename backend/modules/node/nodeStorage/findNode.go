package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"context"
)

func (db *mongoStore) FindNode(ctx context.Context, conditions map[string]interface{}) (*nodeModel.Node, error) {
	collection := db.db.Database("AutomationTest").Collection("Node")

	var data *nodeModel.Node

	if err := collection.FindOne(ctx, conditions).Decode(data); err != nil {
		if err.Error() == common.RecordNotFound {
			return nil, err
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
