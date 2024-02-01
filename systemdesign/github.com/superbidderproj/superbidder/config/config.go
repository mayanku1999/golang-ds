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

var mutex = &sync.Mutex{}

// AppConfig represents the structure of the configuration
type AppConfig struct {
	App      *AppConfigDetails      `yaml:"app"`
	Database *DatabaseConfigDetails `yaml:"database"`
}

// AppConfigDetails represents the details of the 'app' section in the configuration
type AppConfigDetails struct {
	Name      string                    `yaml:"name"`
	Version   string                    `yaml:"version"`
	Debug     bool                      `yaml:"debug"`
	Owners    []string                  `yaml:"owner"`
	MapOwner  []map[string]*MapOwnerVal `yaml:"mapowner"`
	MapOwner1 map[string]*MapOwnerVal   `yaml:"mapowner1"`
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
func LoadSingleInstanceConfigHelper() *AppConfig {
	configOnce.Do(func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		//execRelPath := filepath.Join(filepath.Dir(ex), "..", "..", "superbidderproj", "config", "config.yml")
		//configFileExec, err := ioutil.ReadFile(execRelPath)

		pwd, _ := os.Getwd()
		fmt.Println(pwd)
		configFile, err := ioutil.ReadFile(pwd + "/superbidder/config/config.yml")
		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}

		if err := yaml.Unmarshal(configFile, &config); err != nil {
			log.Fatalf("failed to unmarshal config: %v", err)
		}
	})
	return config
}

func LoadConfig() *AppConfig {
	if config == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if config == nil {
			return LoadSingleInstanceConfigHelper()
		} else {
			fmt.Println("Single Instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single Instance already created-2, returning the same")
	}
	return config
}

//wg := sync.WaitGroup{}
//for _, actor := range actorIds {
//wg.Add(1)
//go func(actorId string) {
//defer wg.Done()
//_, err := s.triggerAndGetPreInvestmentOnboardingDetails(ctx, actorId)
//if err != nil {
//logger.Error(ctx, "error getting pre-investment onboarding details", zap.Error(err), zap.String(logger.ACTOR_ID_V2, actorId))
//return
//}
//}(actor)
//}
//wg.Wait()
