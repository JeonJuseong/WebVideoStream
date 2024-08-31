package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//version: 1.21

type Config struct {
	FilePath string `json:"file_path"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	//videocontrol.CompressVideo("<video absolute path>")
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	portNumber := flag.String("port", "80", "--port=<port number>")
	flag.Parse()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/video/", func(w http.ResponseWriter, r *http.Request) {
		serveVideo(w, r, config.FilePath)
	})
	http.HandleFunc("/websource/", serveWebsource)
	fmt.Println("Starting server at :" + *portNumber)
	http.ListenAndServe("0.0.0.0:"+*portNumber, nil)
}
