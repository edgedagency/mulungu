package configuration

import (
	"os"
	"strings"

	"github.com/creamdog/gonfig"
)

var c *Config

func init() {
	c = NewConfiguration()
}

//Config config sturcture
type Config struct {
	configProvider gonfig.Gonfig
	namespace      string
	path           string
}

//GetConfig returns configuration instance
func GetConfig() *Config {
	return c
}

//NewConfiguration returns app configuration
func NewConfiguration() *Config {
	c = &Config{}

	return c
}

//Load loads configurations from JSON file
func Load() {
	c.Load()
}

//Load loads configurations from JSON file
func (c *Config) Load() error {

	f, fileErr := os.Open(c.path)
	defer f.Close()
	if fileErr != nil {
		return fileErr
	}

	config, configErr := gonfig.FromJson(f)
	if configErr != nil {
		return configErr
	}
	c.configProvider = config

	return nil
}

//AddPath adds configuration file path
func AddPath(path string) { c.AddPath(path) }

//AddPath adds configuration file path
func (c *Config) AddPath(path string) {
	c.path = path //append(c.paths, path)
}

//GetString return string configuration
func GetString(key, defaultValue string) string { return c.GetString(key, defaultValue) }

//GetString return string configuration
func (c *Config) GetString(key, defaultValue string) string {
	key = c.namespaceKey(key)
	config, configErr := c.configProvider.GetString(key, defaultValue)
	if configErr != nil {
		return defaultValue
	}
	return config
}

//Namespace associates a namespace with configurations
func Namespace(namespace string) { c.namespace = namespace }

//Namespace associates a namespace with configurations
func (c *Config) Namespace(namespace string) {
	c.namespace = namespace
}

//namespaceKey returns key namespaced
func (c *Config) namespaceKey(key string) string {
	if c.namespace != "" {
		return strings.Join([]string{c.namespace, key}, "/")
	}
	return key
}
