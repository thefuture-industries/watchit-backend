package packages

import (
	"go-user-service/internal/lib"

	"github.com/noneandundefined/vision-go"
)

type Errors struct {
	monitor *vision.Vision
	logger  *lib.Logger
}

func NewErrors(monitor *vision.Vision) *Errors {
	return &Errors{
		monitor: monitor,
		logger:  lib.NewLogger(),
	}
}

func (h Errors) HandleErrors(err error, isDBError bool) {
	h.logger.Error(err.Error())
	h.monitor.VisionError(err)
	if isDBError {
		h.monitor.VisionDBError()
	}
}
