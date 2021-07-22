package main

import "github.com/fernandodr19/mybank/pkg/instrumentation/logger"

func main() {
	log := logger.Default()
	log.Infoln("=== My Bank API ===")
}
