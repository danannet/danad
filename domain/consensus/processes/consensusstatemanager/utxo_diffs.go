package consensusstatemanager

import (
	"github.com/danannet/danad/domain/consensus/model"
	"github.com/danannet/danad/domain/consensus/model/externalapi"
)

func (csm *consensusStateManager) stageDiff(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	utxoDiff externalapi.UTXODiff, utxoDiffChild *externalapi.DomainHash) {

	log.Tracef("stageDiff start for block %s", blockHash)
	defer log.Tracef("stageDiff end for block %s", blockHash)

	log.Debugf("Staging block %s as the diff child of %s", utxoDiffChild, blockHash)
	csm.utxoDiffStore.Stage(stagingArea, blockHash, utxoDiff, utxoDiffChild)
}
