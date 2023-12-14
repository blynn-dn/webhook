package config

import (
    "fmt"
    yaml "gopkg.in/yaml.v3"
    "log"
    "os"
    "sync"
)

// Config represents the config file
type Config struct {
    Secret   string `yaml:"secret"`
    BasePath string `yaml:"base_path"`
    Port     int    `yaml:"port"`
}

type SafeConfig struct {
    sync.RWMutex
    C *Config
}

// LoadConfig - opens the specified YAML config file
func (sc *SafeConfig) LoadConfig(confFile string) (err error) {
    var c = &Config{}

    yamlReader, err := os.Open(confFile)

    log.Printf("file: %s\n", confFile)

    if err != nil {
        return fmt.Errorf("error reading config file: %s", err)
    }
    defer yamlReader.Close()
    decoder := yaml.NewDecoder(yamlReader)
    decoder.KnownFields(true)

    if err = decoder.Decode(c); err != nil {
        return fmt.Errorf("error parsing config file: %s", err)
    }

    log.Printf("->%v\n", c)

    // lever the mutex to lock the config before updating
    sc.Lock()
    sc.C = c
    sc.Unlock()

    return nil
}
