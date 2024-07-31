package main

import "dsp/proxy_server/tools/log"

func main() {
	
	logger := log.NewLogger(log.DEBUG)

	logger.Debug("started server")

}
