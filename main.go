package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	config := readConfiguration()
	Port := config.Port
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		urlParams := strings.Split(urlPath, "/")
		if len(urlParams) == 3 {
			CollectDataToSend(w, r, urlParams, &config)
		} else {
			w.WriteHeader(500)
		}
	})
	http.ListenAndServe(":" + Port, nil)
}

func readConfiguration() (configuration) {
	pathToFile := os.Getenv("LOGGER_CONFIGURATION_FILE")
	if pathToFile == "" {
		pathToFile = "/home/pi/go/src/go-log-to-elasticsearch/configuration.yaml"
	}
	yamlFile, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		panic(err)
	}

	var config configuration

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Configuration Loaded : %+v \n", config)
	}
	return config
}

func CollectDataToSend(w http.ResponseWriter, r *http.Request, urlParams []string, config *configuration) {
	originName := urlParams[1]
	originType := urlParams[2]
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	postingUrl := config.ElasticsearchURL + "/" + originName + "/" + originType
	_, err := client.Post(postingUrl, "application/json" , r.Body)
	if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}
