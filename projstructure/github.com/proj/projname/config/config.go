package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// AppConfig represents the structure of the configuration
type AppConfig struct {
	App      *AppConfigDetails      `yaml:"app"`
	Database *DatabaseConfigDetails `yaml:"database"`
}

// AppConfigDetails represents the details of the 'app' section in the configuration
type AppConfigDetails struct {
	Name     string                    `yaml:"name"`
	Version  string                    `yaml:"version"`
	Debug    bool                      `yaml:"debug"`
	Owners   []string                  `yaml:"owner"`
	MapOwner []map[string]*MapOwnerVal `yaml:"mapowner"`
}

type MapOwnerVal struct {
	Value1 []string `yaml:"value1"`
	Value2 bool     `yaml:"value2"`
}

// DatabaseConfigDetails represents the details of the 'database' section in the configuration
type DatabaseConfigDetails struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

var (
	config     *AppConfig
	configOnce sync.Once
)

// LoadConfig loads the configuration from the config.yml file
func LoadConfig() *AppConfig {
	configOnce.Do(func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		pwd, _ := os.Getwd()
		fmt.Println(exPath)
		fmt.Println(pwd)

		//execRelPath := filepath.Join(exPath, "..", "..", "projname", "config", "config.yml")
		//configFile, err := ioutil.ReadFile(execRelPath)

		configFile, err := ioutil.ReadFile(pwd + "/projname/config/config.yml")

		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}
		if err := yaml.Unmarshal(configFile, &config); err != nil {
			log.Fatalf("failed to unmarshal config: %v", err)
		}
	})

	return config
}
