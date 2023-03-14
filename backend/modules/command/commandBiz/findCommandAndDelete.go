package commandBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type FindCommandStore interface {
	ListCommand(ctx context.Context, conditions interface{}) ([]commandModel.CommandNode, error)
	DeleteCommand(ctx context.Context, conditions interface{}) error
}

type findCommandBiz struct {
	store FindCommandStore
}

func NewFindCommandBiz(store FindCommandStore) *findCommandBiz {
	return &findCommandBiz{store}
}

func (biz *findCommandBiz) FindCommandAndDelete(ctx context.Context, data *commandModel.Node) ([]commandModel.CommandNode, error) {
	result, err := biz.store.ListCommand(ctx, bson.M{"node_id": data.NodeId})

	if err != nil {
		return nil, common.ErrCannotDeleteEntity("Can not list command", err)
	}

	if err := biz.store.DeleteCommand(ctx, bson.M{"node_id": data.NodeId}); err != nil {
		return nil, common.ErrCannotDeleteEntity("Node Command", err)
	}

	if err != nil {
		return nil, common.ErrCannotDeleteEntity("Node Command", err)
	}
	return result, nil
}
