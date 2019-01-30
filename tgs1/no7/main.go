package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shafwany/Pemrograman-Jaringan/tree/master/tgs1/no7/config"
	"github.com/spf13/viper"
)

/*
var (
	cnf        *Config
	configPath string
)
func initConfig() error {
	flag.StringVar(&configPath, "c", "config.yaml", "Configuration File")
	flag.Parse()
	c, err := NewCfg(configPath)
	if err != nil {
		return err
	}
	cnf = c
	return err
}
*/
func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration
	/*
		err := initConfig()
		if err != nil {
			log.Fatal(err)
		}
	*/
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Printf("Database = %s", configuration.Database.ConnectionUri)
	log.Printf("Port = %d", configuration.Server.Port)
	log.Printf("Directory = %s", configuration.Directory.Dir)

	fs := http.FileServer(http.Dir(configuration.Directory.Dir))
	http.Handle("/", fs)

	log.Println("Listening...")
	port := fmt.Sprintf(":%d", configuration.Server.Port)
	http.ListenAndServe(port, nil)
}