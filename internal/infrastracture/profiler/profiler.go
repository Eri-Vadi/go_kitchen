package profiler

import (
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/logger"
	"net/http"
	_ "net/http/pprof"
)

func StartProfiler() {
	logger.LogMessage("Starting kitchen profiler. Access http://localhost:6060/debug/pprof/")
	go func() {
		logger.LogError(http.ListenAndServe("localhost:6060", nil).Error())
	}()
}