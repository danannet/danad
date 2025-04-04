package transactionvalidator

import (
	"github.com/danannet/danad/domain/consensus/model/externalapi"
)

// PopulateMass calculates and populates the mass of the given transaction
func (v *transactionValidator) PopulateMass(transaction *externalapi.DomainTransaction) {
	if transaction.Mass != 0 {
		return
	}
	transaction.Mass = v.txMassCalculator.CalculateTransactionMass(transaction)
}
