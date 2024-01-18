package antchain

import (
	"context"
	"fmt"
)

func (c *client) QueryTransaction(ctx context.Context, hash string) (string, error) {
	return c.chainCall(ctx, "QUERYTRANSACTION", WithParam("hash", hash))
}

func (c *client) QueryReceipt(ctx context.Context, hash string) (string, error) {
	return c.chainCall(ctx, "QUERYRECEIPT", WithParam("hash", hash))
}

func (c *client) QueryBlockHeader(ctx context.Context, blockNumber int64) (string, error) {
	return c.chainCall(ctx, "QUERYBLOCK", WithParam("requestStr", blockNumber))
}

func (c *client) QueryBlockBody(ctx context.Context, blockNumber int64) (string, error) {
	return c.chainCall(ctx, "QUERYBLOCKBODY", WithParam("requestStr", blockNumber))
}

func (c *client) QueryLastBlock(ctx context.Context) (string, error) {
	return c.chainCall(ctx, "QUERYLASTBLOCK")
}

func (c *client) QueryAccount(ctx context.Context, account string) (string, error) {
	return c.chainCall(ctx, "QUERYACCOUNT", WithParam("requestStr", fmt.Sprintf(`{"queryAccount":"%s"}`, account)))
}
