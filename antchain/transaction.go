package antchain

import "context"

func (c *client) CreateAccount(ctx context.Context, account, kmsID string, gas int) (string, error) {
	return c.chainCallForBiz(ctx, "TENANTCREATEACCUNT",
		WithParam("newAccountId", account),
		WithParam("newAccountKmsId", kmsID),
		WithParam("gas", gas),
	)
}

func (c *client) Deposit(ctx context.Context, content string, gas int) (string, error) {
	return c.chainCallForBiz(ctx, "DEPOSIT",
		WithParam("content", content),
		WithParam("gas", gas),
	)
}

func (c *client) DeploySolidity(ctx context.Context, name, code string, gas int) (string, error) {
	return c.chainCallForBiz(ctx, "DEPLOYCONTRACTFORBIZ",
		WithParam("contractName", name),
		WithParam("contractCode", code),
		WithParam("gas", gas),
	)
}

func (c *client) AsyncCallSolidity(ctx context.Context, contractName, methodSign, inputParams, outTypes string, gas int) (string, error) {
	return c.chainCallForBiz(ctx, "CALLCONTRACTBIZASYNC",
		WithParam("contractName", contractName),
		WithParam("methodSignature", methodSign),
		WithParam("inputParamListStr", inputParams),
		WithParam("outTypes", outTypes),
		WithParam("gas", gas),
	)
}
