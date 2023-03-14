package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddNodeStore interface {
	FindNode(ctx context.Context, conditions map[string]interface{}) (*nodeModel.Node, error)
	CreateNode(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error)
	DeleteNode(ctx context.Context, conditions interface{}) error
}

type addNodeBiz struct {
	store AddNodeStore
}

func NewAddNodeBiz(store AddNodeStore) *addNodeBiz {
	return &addNodeBiz{store}
}

func (biz *addNodeBiz) AddNewNode(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	if err := biz.store.DeleteNode(ctx, bson.M{"node_id": data.NodeId}); err != nil {
		// nay em thay no khong tra ra loi khi tim thay :V, nen loi chac cua database thoi
		return nil, common.ErrDB(err)
	}

	result, err := biz.store.CreateNode(ctx, data)
	if err != nil {
		component.InfoLogger.Println("Can not Create Node")
		return nil, common.ErrCannotCreateEntity("Node Registerd", err)
	}

	return result, nil
}
