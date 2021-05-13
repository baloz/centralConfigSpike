package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func readUsingConsule() {
	viper.AddRemoteProvider("consul", "localhost:8500", "config/configurations.json")
	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("--> ", viper.Get("myname"))
}