package libs

import (
	"os"
	"gopkg.in/yaml.v2"
	"fmt"
)

type Config struct {
	Database struct {
		DBName string `yaml:"dbname"`
	} `yaml:"database"`
}

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}
var cfg Config

func GetConfigValues() Config {
	file, err := os.Open("conf/config.yml")
	if err != nil {
	    processError(err)
	} else {
		defer file.Close()

		
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&cfg)
		if err != nil {
		    processError(err)
		} else {
			return cfg
		}
	} 
	return cfg
}