package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddResultStore interface {
	PostResult(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error)
	DeleteNodeOldResult(ctx context.Context, conditions interface{}) error
}

type addResultBiz struct {
	store AddResultStore
}

func NewAddResultBiz(store AddResultStore) *addResultBiz {
	return &addResultBiz{store}
}

func (biz *addResultBiz) AddNewResult(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error) {

	// xoa ban cu
	if err := biz.store.DeleteNodeOldResult(ctx, bson.M{"node_id": data.NodeId}); err != nil {
		return nil, common.ErrDB(err)
	}
	result, err := biz.store.PostResult(ctx, data)
	if err != nil {
		return nil, common.ErrDB(err)
	}
	// chi can lay ket qua moi nhat thoi

	return result, nil
}
