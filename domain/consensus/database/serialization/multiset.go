package serialization

import (
	"github.com/danannet/danad/domain/consensus/model"
	"github.com/danannet/danad/domain/consensus/utils/multiset"
)

// MultisetToDBMultiset converts Multiset to DbMultiset
func MultisetToDBMultiset(multiset model.Multiset) *DbMultiset {
	return &DbMultiset{
		Multiset: multiset.Serialize(),
	}
}

// DBMultisetToMultiset converts DbMultiset to Multiset
func DBMultisetToMultiset(dbMultiset *DbMultiset) (model.Multiset, error) {
	return multiset.FromBytes(dbMultiset.Multiset)
}
