package services

import "github.com/askaroe/geminiApiGolang/jsonlog"

type Service struct {
	Logger *jsonlog.Logger
}

func NewService(logger *jsonlog.Logger) *Service {
	return &Service{Logger: logger}
}
