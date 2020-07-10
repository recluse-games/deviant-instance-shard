package actions

import (
	"go.uber.org/zap"
)

// GetLogger Returns a zap logger for this package.
func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}
