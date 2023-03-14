package nodeModel

import (
	"errors"
	"strings"
)

type Node struct {
	NodeId   string `json:"node_id" bson:"node_id"`
	UserName string `json:"user_name"bson:"user_name,omitempty"`
	Result   string `json:"result" bson:"result,omitempty"`
}

type Filter struct {
	NodeId string `json:"node_id" bson:"node_id" "`
}

type Result struct {
	Result string `json:"result" bson:"result,omitempty"`
}

func (node *Node) Validate() error {

	// check validate of node name

	node.NodeId = strings.TrimSpace(node.NodeId)

	if node.NodeId == "" {
		return errors.New("username name can not be blank")
	}

	return nil
}
