package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/narharim/go-learning/parse-config/config"
)

// Read from file config.json (config.ReadConfig(filePath))
// Parse the json, xml in structs
// create getters
func main() {

	configPath := flag.String("config", "config.json", "path to the configuration file")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)
}
