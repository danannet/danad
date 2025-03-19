package protowire

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DanadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DanadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *DanadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
