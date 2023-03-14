package commandStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
)

func (db *mongoStore) CreateNodeCommand(ctx context.Context, data *commandModel.CommandNode) error {
	collection := db.db.Database("AutomationTest").Collection("Node_Command")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
