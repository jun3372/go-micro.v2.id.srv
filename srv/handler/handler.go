package handler

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/snowflake"

	"jun.srv.id/proto"
)

var (
	nodes map[int64]*snowflake.Node
	node  *snowflake.Node
	err   error
)

type IdHandler struct {
}

func (i IdHandler) MakeNode(n int64) (*snowflake.Node, error) {
	if n < 1 {
		n = int64(rand.Intn(1022) + 1)
	}

	node, err = snowflake.NewNode(n)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (i IdHandler) GetId(ctx context.Context, request *proto.IdRequest, response *proto.IdResponse) error {
	node, err = i.MakeNode(request.Node)
	if err != nil {
		return err
	}

	response.Id, err = strconv.ParseInt(node.Generate().String(), 10, 64)
	response.Node = request.GetNode()
	if err != nil {
		return err
	}

	return nil
}

func NewIdHandler() *IdHandler {
	return &IdHandler{}
}
