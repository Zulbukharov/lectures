package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Settings domain with application settings variables
type Settings struct {
	PostgresUser     string `yaml:"postgres_user"`
	PostgresPassword string `yaml:"postgres_password"`
	PostgresHost     string `yaml:"postgres_host"`
	PostgresPort     string `yaml:"postgres_port"`
	PostgresDBName   string `yaml:"postgres_db_name"`
	// ServerHTTPPort   string `yaml:"server_http_port"`
}

// ParseYAML parses given yaml file and returns Settings instance
func ParseYAML(fileName string) (Settings, error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Settings{}, err
	}
	settings := Settings{}
	if err := yaml.Unmarshal([]byte(body), &settings); err != nil {
		return settings, err
	}
	return settings, nil
}
