package rpchandlers

import (
	"github.com/danannet/danad/app/appmessage"
	"github.com/danannet/danad/app/rpc/rpccontext"
	"github.com/danannet/danad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
