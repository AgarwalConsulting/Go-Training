package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	appEnv, ok := os.LookupEnv("GO_APP_ENV")
	var err error

	if ok && appEnv == "production" {
		logger, err = zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: true,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build()
		logger.Info("Applying production config...")
	} else {
		// logger, err = zap.NewDevelopment()
		logger, err = zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: true,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build()
		logger.Info("Applying local config...")
	}

	if err != nil {
		fmt.Println("Unable to create logger:", err)

		os.Exit(1)
	}
}

func main() {
	// logger.Info("Log any unstructured data")

	logger.Debug("calling logger.Debug here!")
	logger.Info("calling logger.Info here!")
	logger.Warn("calling logger.Warn here!")
	logger.Error("calling logger.Error here!")
	// logger.Panic("calling logger.Panic here!")
	// logger.Fatal("calling logger.Fatal here!")

	sugar := logger.Sugar()

	sugar.Infow("this calls logger.Info under the hood", "func", "main", "today", "Wednesday")
}

// Log levels
// - Debug
// - Info
// - Warn
// - Error
// - Fatal / Panic
