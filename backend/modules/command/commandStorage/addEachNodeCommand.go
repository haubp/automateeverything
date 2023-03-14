package commandStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
)

func (db *mongoStore) AddEachNodeCommand(ctx context.Context, data *commandModel.EachNodeCommand) error {
	collection := db.db.Database("AutomationTest").Collection("Each_Node_Command")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
