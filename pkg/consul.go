package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func readUsingConsule() {
	viper.AddRemoteProvider("consul", "localhost:8500", "spike")
	viper.SetConfigType("json") // Need to explicitly set this to json
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("--> ", viper.Get("App3"))
}