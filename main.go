package main

import (
	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var log = logging.MustGetLogger("default")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
)


func main() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.NOTICE, "")
	logging.SetBackend(backendLeveled, backendFormatter)

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
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	postingUrl := config.ElasticsearchURL + "/" + originName + "/" + originType
	_, err := client.Post(postingUrl, "application/json" , r.Body)
	if err != nil {
		w.WriteHeader(500)
	} else {
		r.Body.Close()
		w.WriteHeader(200)
	}
}
