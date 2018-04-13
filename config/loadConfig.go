package config

import (
	"log"

	yaml "gopkg.in/yaml.v1"
)

var (
	C       map[string]string
	configs ConfigPage
)

type ConfigPage map[string](map[string]string)

func init() {
	configs = make(ConfigPage)
	err := yaml.Unmarshal([]byte(yamlFile), &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

// LoadConfig by env
func LoadConfig(env string) {
	C = configs[env]
}
