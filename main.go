package main

import (
	config "aqaryint/aqaryint/src/configs"
	router "aqaryint/aqaryint/src/router"
)

func main() {

	config.Init()
	config.Appconfig = config.GetConfig()
	router.Init()

}
