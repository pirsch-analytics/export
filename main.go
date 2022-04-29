package main

import (
	"github.com/BurntSushi/toml"
	exp "github.com/pirsch-analytics/export/exports"
	"github.com/pirsch-analytics/pirsch-go-sdk"
	"log"
	"os"
	"strings"
	"time"
)

var (
	exports = []string{
		"conversion_goals_day",
	}

	config Config
	client *pirsch.Client
)

type Config struct {
	ClientID     string    `toml:"client_id"`
	ClientSecret string    `toml:"client_secret"`
	Hostname     string    `toml:"hostname"`
	Export       []string  `toml:"export"`
	From         time.Time `toml:"from"`
	To           time.Time `toml:"to"`
}

func loadConfig() error {
	data, err := os.ReadFile("config.toml")

	if err != nil {
		return err
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}

func runExports() error {
	for _, export := range config.Export {
		switch export {
		case "conversion_goals_day":
			if err := exp.ExportConversionGoalsDays(client, config.From, config.To); err != nil {
				return err
			}
		}
	}

	return nil
}

func isValidExport(export string) bool {
	for _, e := range exports {
		if export == e {
			return true
		}
	}

	return false
}

func main() {
	if err := loadConfig(); err != nil {
		log.Printf("Error loading config.toml: %v", err)
		return
	}

	if config.ClientID == "" || config.ClientSecret == "" {
		log.Println("Please configure a client ID and secret (created on the developer settings page on pirsch.io)")
		return
	}

	if len(config.Export) == 0 {
		log.Printf("Please specify which exports you would like to run. Available: %s", strings.Join(exports, ", "))
		return
	}

	for _, export := range config.Export {
		if !isValidExport(export) {
			log.Printf("The export option '%s' does not exist. Available: %s", export, strings.Join(exports, ", "))
			return
		}
	}

	if err := os.Mkdir("export", 0744); err != nil && !os.IsExist(err) {
		log.Printf("Error creating export directory: %v", err)
		return
	}

	client = pirsch.NewClient(config.ClientID, config.ClientSecret, config.Hostname, nil)

	if err := runExports(); err != nil {
		log.Printf("Error exporting statistics: %v", err)
		return
	}

	log.Println("Export successful!")
}
