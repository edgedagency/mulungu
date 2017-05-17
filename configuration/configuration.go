package configuration

import (
	"os"
	"strings"

	"github.com/creamdog/gonfig"
	"github.com/spf13/cast"
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

//Get returns an interface. For a specific value use one of the Get methods.
func Get(key string) interface{} { return c.Get(key) }

//Get returns an interface. For a specific value use one of the Get methods.
func (c *Config) Get(key string) interface{} {
	lcaseKey := strings.ToLower(key)
	val := v.find(lcaseKey)
	if val == nil {
		return nil
	}
}

//NewConfiguration returns app configuration
func NewConfiguration() *Config {
	c = &Config{}

	return c
}

//Load loads configurations from JSON file
// func Load() {
// 	c.Load()
// }
//Load loads configurations from config file
func Load(filename string, dest map[string]string) {
	c.Load(filename, dest)

}

//Load adds or updates entries in an existing map with key and values
func (c *Config) Load(filename string, dest map[string]string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	f, fileErr := os.Open(filename)
	//defer f.Close()
	if fileErr != nil {
		return fileErr
	}
	buff := make([]byte, fi.Size())
	f.Read(buff)
	f.Close()

	//
	// config, configErr := gonfig.FromJson(f)
	// if configErr != nil {
	// 	return configErr
	// }
	// c.configProvider = config

	return nil
}

//AddPath adds configuration file path
func AddPath(path string) { c.AddPath(path) }

//AddPath adds configuration file path
func (c *Config) AddPath(path string) {
	c.path = path //append(c.paths, path)

}

//GetString return string configuration
func GetString(key string, defaultValue string) string { return c.GetString(key, defaultValue) }

//GetString return string configuration **593
func (c *Config) GetString(key string, defaultValue string) string {
	key = c.namespaceKey(key)
	config, configErr := c.configProvider.GetString(key, defaultValue)
	if configErr != nil {
		return defaultValue
	}
	return config
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 { return c.GetInt64(key) }

// GetInt64 returns the value associated with the key as an integer.
func (c *Config) GetInt64(key string) int64 {
	return cast.ToInt64(c.Get(key))
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
