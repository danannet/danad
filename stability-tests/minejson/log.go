package main

import (
	"github.com/danannet/danad/infrastructure/logger"
	"github.com/danannet/danad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MNJS")
	spawn      = panics.GoroutineWrapperFunc(log)
)
