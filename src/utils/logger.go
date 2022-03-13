package utils

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}