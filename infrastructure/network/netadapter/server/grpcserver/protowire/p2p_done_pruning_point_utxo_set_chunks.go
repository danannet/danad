package protowire

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DanadMessage_DonePruningPointUtxoSetChunks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DanadMessage_DonePruningPointUtxoSetChunks is nil")
	}
	return &appmessage.MsgDonePruningPointUTXOSetChunks{}, nil
}

func (x *DanadMessage_DonePruningPointUtxoSetChunks) fromAppMessage(_ *appmessage.MsgDonePruningPointUTXOSetChunks) error {
	x.DonePruningPointUtxoSetChunks = &DonePruningPointUtxoSetChunksMessage{}
	return nil
}
