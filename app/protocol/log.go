package protocol

import (
	"github.com/danannet/danad/infrastructure/logger"
	"github.com/danannet/danad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
