package model

import "github.com/danannet/danad/domain/consensus/model/externalapi"

// FinalityManager provides method to validate that a block does not violate finality
type FinalityManager interface {
	VirtualFinalityPoint(stagingArea *StagingArea) (*externalapi.DomainHash, error)
	FinalityPoint(stagingArea *StagingArea, blockHash *externalapi.DomainHash, isBlockWithTrustedData bool) (*externalapi.DomainHash, error)
}
