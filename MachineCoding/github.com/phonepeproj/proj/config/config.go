package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	config     *AppConfig
	configOnce sync.Once
	mutex      = &sync.Mutex{}
)

// AppConfig represents the structure of the configuration
type AppConfig struct {
	App               *AppConfigDetails                `yaml:"App"`
	FeatureConstraint map[string]*UserGroupConstraints `yaml:"FeatureConstraint"`
}

type UserGroupConstraints struct {
	UserGroupConstraint map[string]bool `yaml:"UserGroupConstraint"`
}

// AppConfigDetails represents the details of the 'app' section in the configuration
type AppConfigDetails struct {
	Name string `yaml:"Name"`
}

// LoadConfig loads the configuration from the config.yml file
func LoadConfigHelper() *AppConfig {
	configOnce.Do(func() {

		//ex, err := os.Executable()
		//if err != nil {
		//	log.Panicf("failed to get os path:  %v", err)
		//}
		//execTargetPath := filepath.Dir(ex)
		//execRelPath := filepath.Join(execTargetPath, "..", "..", "/proj", "config", "config.yml")
		//configExecFile, configExecFileErr := ioutil.ReadFile(execRelPath)
		//if configExecFileErr != nil {
		//	log.Panicf("failed to load exec config : %v", configExecFileErr)
		//}

		pwd, _ := os.Getwd()
		configFile, err := ioutil.ReadFile(pwd + "/proj/config/config.yml")
		if err != nil {
			log.Panicf("failed to read config file: %v", err)
		}

		if err := yaml.Unmarshal(configFile, &config); err != nil {
			log.Panicf("failed to unmarshal config: %v", err)
		}
	})

	return config
}

// GetSingleConfigInstance will return the singleton class for app config that will get initialised only once,
// taken a lock on the global config if two or more threads are trying to initialise at the same time.
func GetSingleConfigInstance() *AppConfig {
	if config == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if config == nil {
			return LoadConfigHelper()
		} else {
			fmt.Println("Single Instance already created")
		}
	} else {
		fmt.Println("Single Instance already created")
	}
	return config
}
