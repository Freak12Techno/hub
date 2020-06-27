package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func QueryPlan(ctx context.CLIContext, id uint64) (*types.Plan, error) {
	params := types.NewQueryPlanParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlan)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plan found")
	}

	var plan types.Plan
	if err := ctx.Codec.UnmarshalJSON(res, &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func QueryPlans(ctx context.CLIContext) (types.Plans, error) {
	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlans)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plans found")
	}

	var plans types.Plans
	if err := ctx.Codec.UnmarshalJSON(res, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func QueryPlansForProvider(ctx context.CLIContext, address hub.ProvAddress) (types.Plans, error) {
	params := types.NewQueryPlansForProviderParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlansForProvider)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plans found")
	}

	var plans types.Plans
	if err := ctx.Codec.UnmarshalJSON(res, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func QuerySubscription(ctx context.CLIContext, id uint64) (*types.Subscription, error) {
	params := types.NewQuerySubscriptionParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscription found")
	}

	var subscription types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscription); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func QuerySubscriptions(ctx context.CLIContext) (types.Subscriptions, error) {
	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptions)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForAddress(ctx context.CLIContext, address sdk.AccAddress) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForAddressParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForPlan(ctx context.CLIContext, id uint64) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForPlanParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForPlan)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForNode(ctx context.CLIContext, address hub.NodeAddress) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForNodeParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QueryMembersForSubscription(ctx context.CLIContext, id uint64) ([]sdk.AccAddress, error) {
	params := types.NewQueryMembersForSubscriptionParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryMembersForSubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var members []sdk.AccAddress
	if err := ctx.Codec.UnmarshalJSON(res, &members); err != nil {
		return nil, err
	}

	return members, nil
}
