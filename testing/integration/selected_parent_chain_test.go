package integration

import (
	"testing"

	"github.com/danannet/danad/domain/consensus/model/externalapi"

	"github.com/danannet/danad/app/appmessage"
	"github.com/danannet/danad/domain/consensus/utils/consensushashing"
)

func TestVirtualSelectedParentChain(t *testing.T) {
	// Setup a couple of danad instances
	danad1, danad2, _, teardown := standardSetup(t)
	defer teardown()

	// Register to virtual selected parent chain changes
	onVirtualSelectedParentChainChangedChan := make(chan *appmessage.VirtualSelectedParentChainChangedNotificationMessage)
	err := danad1.rpcClient.RegisterForVirtualSelectedParentChainChangedNotifications(true,
		func(notification *appmessage.VirtualSelectedParentChainChangedNotificationMessage) {
			onVirtualSelectedParentChainChangedChan <- notification
		})
	if err != nil {
		t.Fatalf("Failed to register for virtual selected parent chain change notifications: %s", err)
	}

	// In danad1, mine a chain over the genesis and make sure
	// each chain changed notifications contains only one entry
	// in `added` and nothing in `removed`
	chain1TipHash := consensushashing.BlockHash(danad1.config.NetParams().GenesisBlock)
	chain1TipHashString := chain1TipHash.String()
	const blockAmountToMine = 10
	for i := 0; i < blockAmountToMine; i++ {
		minedBlock := mineNextBlock(t, danad1)
		notification := <-onVirtualSelectedParentChainChangedChan
		if len(notification.RemovedChainBlockHashes) > 0 {
			t.Fatalf("RemovedChainBlockHashes is unexpectedly not empty")
		}
		if len(notification.AddedChainBlockHashes) != 1 {
			t.Fatalf("Unexpected length of AddedChainBlockHashes. Want: %d, got: %d",
				1, len(notification.AddedChainBlockHashes))
		}

		minedBlockHash := consensushashing.BlockHash(minedBlock)
		minedBlockHashString := minedBlockHash.String()
		if minedBlockHashString != notification.AddedChainBlockHashes[0] {
			t.Fatalf("Unexpected block hash in AddedChainBlockHashes. Want: %s, got: %s",
				minedBlockHashString, notification.AddedChainBlockHashes[0])
		}
		chain1TipHashString = minedBlockHashString
	}

	// In danad2, mine a different chain of `blockAmountToMine` + 1
	// blocks over the genesis
	var chain2Tip *externalapi.DomainBlock
	for i := 0; i < blockAmountToMine+1; i++ {
		chain2Tip = mineNextBlock(t, danad2)
	}

	// Connect the two danads. This should trigger sync
	// between the two nodes
	connect(t, danad1, danad2)

	chain2TipHash := consensushashing.BlockHash(chain2Tip)
	chain2TipHashString := chain2TipHash.String()

	// For the first `blockAmountToMine - 1` blocks we don't expect
	// the chain to change at all, thus there will be no notifications

	// Either the next block or the one after it will cause a reorg
	reorgNotification := <-onVirtualSelectedParentChainChangedChan

	// Make sure that the reorg notification contains exactly
	// `blockAmountToMine` blocks in its `removed`
	if len(reorgNotification.RemovedChainBlockHashes) != blockAmountToMine {
		t.Fatalf("Unexpected length of reorgNotification.RemovedChainBlockHashes. Want: %d, got: %d",
			blockAmountToMine, len(reorgNotification.RemovedChainBlockHashes))
	}

	// Get the virtual selected parent chain from the tip of
	// the first chain
	virtualSelectedParentChainFromChain1Tip, err := danad1.rpcClient.GetVirtualSelectedParentChainFromBlock(
		chain1TipHashString, true)
	if err != nil {
		t.Fatalf("GetVirtualSelectedParentChainFromBlock failed: %s", err)
	}

	// Make sure that `blockAmountToMine` blocks were removed
	// and `blockAmountToMine + 1` blocks were added
	if len(virtualSelectedParentChainFromChain1Tip.RemovedChainBlockHashes) != blockAmountToMine {
		t.Fatalf("Unexpected length of virtualSelectedParentChainFromChain1Tip.RemovedChainBlockHashes. Want: %d, got: %d",
			blockAmountToMine, len(virtualSelectedParentChainFromChain1Tip.RemovedChainBlockHashes))
	}
	if len(virtualSelectedParentChainFromChain1Tip.AddedChainBlockHashes) != blockAmountToMine+1 {
		t.Fatalf("Unexpected length of virtualSelectedParentChainFromChain1Tip.AddedChainBlockHashes. Want: %d, got: %d",
			blockAmountToMine+1, len(virtualSelectedParentChainFromChain1Tip.AddedChainBlockHashes))
	}

	// Make sure that the last block in `added` is the tip
	// of chain2
	lastAddedChainBlock := virtualSelectedParentChainFromChain1Tip.AddedChainBlockHashes[len(virtualSelectedParentChainFromChain1Tip.AddedChainBlockHashes)-1]
	if lastAddedChainBlock != chain2TipHashString {
		t.Fatalf("Unexpected last added chain block. Want: %s, got: %s",
			chain2TipHashString, lastAddedChainBlock)
	}
}
