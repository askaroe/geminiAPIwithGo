package controllers

import (
	"github.com/askaroe/geminiApiGolang/jsonlog"
	"github.com/askaroe/geminiApiGolang/services"
)

type Controller struct {
	s      *services.Service
	logger *jsonlog.Logger
}

func NewController(s *services.Service, logger *jsonlog.Logger) *Controller {
	return &Controller{s: s, logger: logger}
}
