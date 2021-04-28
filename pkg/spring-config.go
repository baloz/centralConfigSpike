package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func readUsingSpringConfig()() {
	body, err := fetchconfiguration()
	if err != nil {
			panic("couldn't load configuration, cannot start. terminating. error: " + err.Error())
	}
	parseconfiguration(body)
}

func fetchconfiguration() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8500/spike/dev", nil)
	if err != nil {
		log.Print("Error ", err)
	}
	req.SetBasicAuth("admin", "admin")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
			panic("couldn't load configuration, cannot start. terminating. error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func parseconfiguration(body []byte) {
	var cloudconfig Springcloudconfig
	err := json.Unmarshal(body, &cloudconfig)
	if err != nil {
			panic("cannot parse configuration, message: " + err.Error())
	}

	for key, value := range cloudconfig.PropertySources[0].Source {
			viper.Set(key, value)
	}

	log.Print(viper.Get("App1"))
}

type Springcloudconfig struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	PropertySources []Propertysource `json:"propertySources"`
}

type Propertysource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}