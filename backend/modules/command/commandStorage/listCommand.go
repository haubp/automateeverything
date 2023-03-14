package commandStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
)

func (db *mongoStore) ListCommand(ctx context.Context, conditions interface{}) ([]commandModel.CommandNode, error) {

	collection := db.db.Database("AutomationTest").Collection("Each_Node_Command")
	cur, err := collection.Find(ctx, conditions)
	if err != nil {
		return nil, common.ErrDB(err)
	}

	var result []commandModel.CommandNode

	if err = cur.All(ctx, &result); err != nil {
		panic(err)
	}

	return result, nil
}
