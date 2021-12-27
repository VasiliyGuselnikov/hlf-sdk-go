package system

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"github.com/hyperledger/fabric/core/chaincode/lifecycle"
	"github.com/hyperledger/fabric/msp"

	"github.com/s7techlab/hlf-sdk-go/v2/api"
	peerSDK "github.com/s7techlab/hlf-sdk-go/v2/peer"
)

// NewLifecycle returns an implementation of api.Lifecycle interface
func NewLifecycle(peerPool api.PeerPool, identity msp.SigningIdentity) api.Lifecycle {
	return &lifecycleCC{peerPool: peerPool, identity: identity, processor: peerSDK.NewProcessor(``)}
}

var _ api.Lifecycle = (*lifecycleCC)(nil)

type lifecycleCC struct {
	peerPool  api.PeerPool
	identity  msp.SigningIdentity
	processor api.PeerProcessor
}

// QueryInstalledChaincodes returns installed chaincodes list
func (c *lifecycleCC) QueryInstalledChaincodes(ctx context.Context) (*lb.QueryInstalledChaincodesResult, error) {
	resp, err := c.endorse(ctx, ``, lifecycle.QueryInstalledChaincodesFuncName, []byte(``))
	if err != nil {
		return nil, err
	}
	ccData := new(lb.QueryInstalledChaincodesResult)
	if err = proto.Unmarshal(resp, ccData); err != nil {
		return nil, fmt.Errorf(`failed to unmarshal protobuf: %w`, err)
	}
	return ccData, nil
}

// InstallChaincode install chaincode on a peer
func (c *lifecycleCC) InstallChaincode(ctx context.Context, installArgs *lb.InstallChaincodeArgs) (*lb.InstallChaincodeResult, error) {
	var (
		args []byte
		resp []byte
		err  error
	)
	if args, err = proto.Marshal(installArgs); err != nil {
		return nil, fmt.Errorf("marshal args: %w", err)
	}
	resp, err = c.endorse(ctx, ``, lifecycle.InstallChaincodeFuncName, args)
	if err != nil {
		return nil, err
	}

	ccResult := new(lb.InstallChaincodeResult)
	if err = proto.Unmarshal(resp, ccResult); err != nil {
		return nil, fmt.Errorf("unmarshal protobuf: %w", err)
	}

	return ccResult, nil
}

// ApproveFromMyOrg approves chaincode package on a channel
func (c *lifecycleCC) ApproveFromMyOrg(
	ctx context.Context,
	channel api.Channel,
	approveArgs *lb.ApproveChaincodeDefinitionForMyOrgArgs) error {
	var (
		args []byte
		cc   api.Chaincode
		err  error
	)
	if args, err = proto.Marshal(approveArgs); err != nil {
		return fmt.Errorf("marshal args: %w", err)
	}
	cc, err = channel.Chaincode(ctx, lifecycleName)
	if err != nil {
		return fmt.Errorf("initalize chaincode: %w", err)
	}

	_, _, err = cc.Invoke(lifecycle.ApproveChaincodeDefinitionForMyOrgFuncName).
		WithIdentity(c.identity).
		ArgBytes([][]byte{args}).
		Do(ctx, api.WithEndorsingMpsIDs([]string{c.identity.GetMSPIdentifier()}))

	if err != nil {
		return fmt.Errorf("invoke chaincode: %w", err)
	}

	return nil
}

// CheckCommitReadiness returns commitments statuses of participants on chaincode definition
func (c *lifecycleCC) CheckCommitReadiness(ctx context.Context, channelID string, args *lb.CheckCommitReadinessArgs) (
	*lb.CheckCommitReadinessResult, error) {
	var (
		argsBytes []byte
		resp      []byte
		err       error
	)
	if argsBytes, err = proto.Marshal(args); err != nil {
		return nil, fmt.Errorf("marshal args: %w", err)
	}
	resp, err = c.endorse(ctx, channelID, lifecycle.CheckCommitReadinessFuncName, argsBytes)
	if err != nil {
		return nil, fmt.Errorf(`failed to endorse proposal: %w`, err)
	}
	result := new(lb.CheckCommitReadinessResult)
	if err = proto.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal proposal response: %w", err)
	}

	return result, nil
}

func (c *lifecycleCC) endorse(ctx context.Context, channel string, fn string, args ...[]byte) ([]byte, error) {
	processor := peerSDK.NewProcessor(channel)
	prop, _, err := processor.CreateProposal(lifecycleName, c.identity, fn, args, nil)
	if err != nil {
		return nil, fmt.Errorf(`failed to create proposal: %w`, err)
	}

	resp, err := c.peerPool.Process(ctx, c.identity.GetMSPIdentifier(), prop)
	if err != nil {
		return nil, fmt.Errorf(`failed to endorse proposal: %w`, err)
	}
	return resp.Response.Payload, nil
}
