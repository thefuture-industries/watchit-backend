package packages

import (
	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Errors struct {
	monitor *vision.Vision
	logger  *zap.Logger
}

func NewErrors(monitor *vision.Vision, logger *zap.Logger) *Errors {
	return &Errors{
		monitor: monitor,
		logger:  logger,
	}
}

func (h Errors) HandleErrors(err error, isDBError bool) {
	ErrorLog(err)
	h.monitor.VisionError(err)
	if isDBError {
		h.monitor.VisionDBError()
	}
	h.logger.Error(err.Error(), zap.Error(err))
}
