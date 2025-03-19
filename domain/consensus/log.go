package consensus

import (
	"github.com/danannet/danad/infrastructure/logger"
	"github.com/danannet/danad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
