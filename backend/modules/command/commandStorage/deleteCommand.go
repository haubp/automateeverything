package commandStorage

import (
	"backend_autotest/common"
	"context"
)

func (db *mongoStore) DeleteCommand(ctx context.Context, conditions interface{}) error {

	collection := db.db.Database("AutomationTest").Collection("Each_Node_Command")
	_, err := collection.DeleteMany(ctx, conditions)
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
