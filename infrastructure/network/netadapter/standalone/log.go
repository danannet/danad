package standalone

import (
	"github.com/danannet/danad/infrastructure/logger"
	"github.com/danannet/danad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
