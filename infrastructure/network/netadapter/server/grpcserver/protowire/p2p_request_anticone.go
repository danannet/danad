package protowire

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DanadMessage_RequestAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DanadMessage_RequestAnticone is nil")
	}
	return x.RequestAnticone.toAppMessage()
}

func (x *RequestAnticoneMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestAnticoneMessage is nil")
	}
	blockHash, err := x.BlockHash.toDomain()
	if err != nil {
		return nil, err
	}

	contextHash, err := x.ContextHash.toDomain()
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgRequestAnticone{
		BlockHash:   blockHash,
		ContextHash: contextHash,
	}, nil

}

func (x *DanadMessage_RequestAnticone) fromAppMessage(msgRequestPastDiff *appmessage.MsgRequestAnticone) error {
	x.RequestAnticone = &RequestAnticoneMessage{
		BlockHash:   domainHashToProto(msgRequestPastDiff.BlockHash),
		ContextHash: domainHashToProto(msgRequestPastDiff.ContextHash),
	}
	return nil
}
