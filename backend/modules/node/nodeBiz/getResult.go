package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type GetResultStore interface {
	GetResult(ctx context.Context, conditions interface{}) (*nodeModel.Result, error)
}

type getResultBiz struct {
	store GetResultStore
}

func NewGetResultBiz(store GetResultStore) *getResultBiz {
	return &getResultBiz{store: store}
}

func (biz *getResultBiz) GetResult(ctx context.Context, conditions interface{}) (*nodeModel.Result, error) {

	result, err := biz.store.GetResult(ctx, bson.M{"node_id": conditions})
	if err != nil {
		if err.Error() == common.RecordNotFound {
			return nil, common.ErrEntityNotExist("node", nil)
		}
		common.ErrDB(err)
	}
	return result, nil
}
