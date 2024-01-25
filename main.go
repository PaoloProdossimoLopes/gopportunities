package main

import (
	"github.com/PaoloProdossimoLopes/gopportunities/config"
	"github.com/PaoloProdossimoLopes/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	configError := config.Init()
	if configError != nil {
		logger.Errorf("Failure to config application: %v", configError)
		panic(configError)
	}

	router.Initialize()
}
