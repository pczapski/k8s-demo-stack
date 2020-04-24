package logger

import "go.uber.org/zap"

func NewZap(appVersion string) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	logger.With(zap.String("version", appVersion))
	zap.ReplaceGlobals(logger)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
