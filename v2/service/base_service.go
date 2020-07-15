package service

import "github.com/felippeduran/pitaya/v2/pipeline"

type baseService struct {
	handlerHooks *pipeline.HandlerHooks
}

func (h *baseService) SetHandlerHooks(handlerHooks *pipeline.HandlerHooks) {
	h.handlerHooks = handlerHooks
}
