package blockrelay

import (
	"github.com/danannet/danad/app/appmessage"
	peerpkg "github.com/danannet/danad/app/protocol/peer"
	"github.com/danannet/danad/domain"
	"github.com/danannet/danad/infrastructure/network/netadapter/router"
)

// PruningPointProofRequestsContext is the interface for the context needed for the HandlePruningPointProofRequests flow.
type PruningPointProofRequestsContext interface {
	Domain() domain.Domain
}

// HandlePruningPointProofRequests listens to appmessage.MsgRequestPruningPointProof messages and sends
// the pruning point proof to the requesting peer.
func HandlePruningPointProofRequests(context PruningPointProofRequestsContext, incomingRoute *router.Route,
	outgoingRoute *router.Route, peer *peerpkg.Peer) error {

	for {
		_, err := incomingRoute.Dequeue()
		if err != nil {
			return err
		}

		log.Debugf("Got request for pruning point proof from %s", peer)

		pruningPointProof, err := context.Domain().Consensus().BuildPruningPointProof()
		if err != nil {
			return err
		}
		pruningPointProofMessage := appmessage.DomainPruningPointProofToMsgPruningPointProof(pruningPointProof)
		err = outgoingRoute.Enqueue(pruningPointProofMessage)
		if err != nil {
			return err
		}

		log.Debugf("Sent pruning point proof to %s", peer)
	}
}
