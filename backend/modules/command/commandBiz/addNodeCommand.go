package commandBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
)

type NodeCommandStore interface {
	CreateNodeCommand(ctx context.Context, data *commandModel.CommandNode) error
	AddEachNodeCommand(ctx context.Context, data *commandModel.EachNodeCommand) error
}

type nodeCommandBiz struct {
	store NodeCommandStore
}

func NewNodeCommandBiz(store NodeCommandStore) *nodeCommandBiz {
	return &nodeCommandBiz{store}
}

func (biz *nodeCommandBiz) AddNodeCommand(ctx context.Context, data *commandModel.CommandNode) error {

	for _, eachNode := range data.Nodes {
		tmp := commandModel.EachNodeCommand{eachNode.NodeId, data.Command, data.User}
		err := biz.store.AddEachNodeCommand(ctx, &tmp)
		if err != nil {
			return common.ErrCannotCreateEntity("Node Command", err)
		}
	}

	return nil
}
