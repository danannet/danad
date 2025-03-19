package protowire

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DanadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DanadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *DanadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
