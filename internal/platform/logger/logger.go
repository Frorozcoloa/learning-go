package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger inicializa zap de forma global
func InitLogger() {
	// "NewDevelopment" muestra logs bonitos y legibles para ti
	// En producción usaríamos "NewProduction" para JSON
	var err error
	Log, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}