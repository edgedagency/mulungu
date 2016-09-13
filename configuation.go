package mulungu

import (
	"fmt"

	"github.com/edgedagency/mulungu/util"
)

// Configuration is used to
type Configuration struct {
	configurations *map[string]interface{}
}

//NewConfiguration instantiates configuration structure and loads configurations
func NewConfiguration(basePath string) *Configuration {

	c := new(Configuration)
	data, err := util.FileRead(basePath + "conf.json")

	if err != nil {
		panic("Unable to read or reach conf.json err:" + err.Error())
	}
	fmt.Println(string(data))
	configurations, err := util.JSONDecode(data)

	if err != nil {
		panic("Unable to read and configure configurations : " + err.Error())
	}

	c.configurations = &configurations

	fmt.Println(c.configurations)
	return c
}

// Get returns configuration information by key
func (c *Configuration) Get(key string) interface{} {

	return nil
}

// Set creates a configuration item using key and value provided
func (c *Configuration) Set(key string, value interface{}) interface{} {

	return nil
}

// Remove removes item from configuration object by key
func (c *Configuration) Remove(key string) interface{} {

	return nil
}

// Has checks if a configuration item exists by key
func (c *Configuration) Has(key string) bool {

	return true
}
