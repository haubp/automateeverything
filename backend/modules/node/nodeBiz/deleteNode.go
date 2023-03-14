package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteNodeStore interface {
	DeleteNode(ctx context.Context, conditions interface{}) error
}

type deleteNodeBiz struct {
	store DeleteNodeStore
}

func NewDeleteNodeBiz(store DeleteNodeStore) *deleteNodeBiz {
	return &deleteNodeBiz{store}
}

func (biz *deleteNodeBiz) DeleteNode(ctx context.Context, data *nodeModel.Node) error {
	if err := biz.store.DeleteNode(ctx, bson.M{"node_id": data.NodeId}); err != nil {
		// loi do db tra ra
		return common.ErrCannotDeleteEntity("Node", err)
	}
	return nil
}
