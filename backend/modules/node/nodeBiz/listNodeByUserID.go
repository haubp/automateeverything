package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/node/nodeModel"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ListNodeStore interface {
	ListNode(ctx context.Context, conditions map[string]interface{}) ([]nodeModel.Node, error)
}

type listNodeBiz struct {
	store ListNodeStore
}

func NewListNodeBiz(store ListNodeStore) *listNodeBiz {
	return &listNodeBiz{store}
}

func (biz *listNodeBiz) ListNode(ctx context.Context, data *userModel.User) ([]nodeModel.Node, error) {
	result, err := biz.store.ListNode(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		return nil, common.ErrCannotListEntity("Node", err)
	}

	return result, nil

}
