package protowire

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/danannet/danad/infrastructure/network/netadapter/id"
	"github.com/danannet/danad/util/mstime"
	"github.com/pkg/errors"
)

func (x *DanadMessage_Version) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DanadMessage_Version is nil")
	}
	return x.Version.toAppMessage()
}

func (x *VersionMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "VersionMessage is nil")
	}
	address, err := x.Address.toAppMessage()
	// Address is optional for non-listening nodes
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	subnetworkID, err := x.SubnetworkId.toDomain()
	//  Full kaspa nodes set SubnetworkId==nil
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	err = appmessage.ValidateUserAgent(x.UserAgent)
	if err != nil {
		return nil, err
	}

	if x.Id == nil {
		return nil, errors.Wrapf(errorNil, "VersionMessage.Id is nil")
	}

	appMsgID, err := id.FromBytes(x.Id)
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgVersion{
		ProtocolVersion: x.ProtocolVersion,
		Network:         x.Network,
		Services:        appmessage.ServiceFlag(x.Services),
		Timestamp:       mstime.UnixMilliseconds(x.Timestamp),
		Address:         address,
		ID:              appMsgID,
		UserAgent:       x.UserAgent,
		DisableRelayTx:  x.DisableRelayTx,
		SubnetworkID:    subnetworkID,
	}, nil
}

func (x *DanadMessage_Version) fromAppMessage(msgVersion *appmessage.MsgVersion) error {
	err := appmessage.ValidateUserAgent(msgVersion.UserAgent)
	if err != nil {
		return err
	}

	versionID, err := msgVersion.ID.SerializeToBytes()
	if err != nil {
		return err
	}

	// Address is optional for non-listening nodes
	var address *NetAddress
	if msgVersion.Address != nil {
		address = appMessageNetAddressToProto(msgVersion.Address)
	}

	x.Version = &VersionMessage{
		ProtocolVersion: msgVersion.ProtocolVersion,
		Network:         msgVersion.Network,
		Services:        uint64(msgVersion.Services),
		Timestamp:       msgVersion.Timestamp.UnixMilliseconds(),
		Address:         address,
		Id:              versionID,
		UserAgent:       msgVersion.UserAgent,
		DisableRelayTx:  msgVersion.DisableRelayTx,
		SubnetworkId:    domainSubnetworkIDToProto(msgVersion.SubnetworkID),
	}
	return nil
}
