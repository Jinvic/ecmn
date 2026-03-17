package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Webhook WebhookConfig `yaml:"webhook"`
	Logging LoggingConfig `yaml:"logging"`
	SMTP    SMTPConfig    `yaml:"smtp"`
	API     APIConfig     `yaml:"api"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type WebhookConfig struct {
	Secret string `yaml:"secret"`
	Path   string `yaml:"path"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type SMTPConfig struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	From     string   `yaml:"from"`
	To       []string `yaml:"to"`
}

type APIConfig struct {
	BaseURL string `yaml:"base_url"`
	Token   string `yaml:"token"`
	Timeout int    `yaml:"timeout"`
}

var AppConfig *Config

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	AppConfig = &cfg
	return &cfg, nil
}

func Get() *Config {
	return AppConfig
}
