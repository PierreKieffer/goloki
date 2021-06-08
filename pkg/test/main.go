package main

import (
	"encoding/json"
	"fmt"
	"github.com/PierreKieffer/appConfig"
	"goloki/pkg/configuration"
	"goloki/pkg/middleware"
	"os"
)

func main() {
	configuration := configuration.Configuration{}
	appConfig.InitConfig(os.Args[1], &configuration)

	var labels = make(map[string]interface{})
	labels["level"] = "INFO"
	labels["foo"] = "bar"
	log := middleware.InitLog("This is an interesting log", labels)

	fmt.Println(log)

}
