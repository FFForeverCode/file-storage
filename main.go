package main

import (
	"file-storage/logger"
)

func main() {

	if err := logger.InitLogger(); err != nil {
		panic("init logger failed: " + err.Error())
	}
	defer logger.Sync()

	logger.Logger.Info("Service started")
}
