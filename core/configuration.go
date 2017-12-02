package core

import "golang.org/x/net/context"

var configuration *Configuration

func init() {
	configuration = &Configuration{}
}

//Configuration used to manage and store system services configurations
type Configuration struct {
}

//Get return based on pointer variable
func Get(ctx context.Context, namespace string, key string, defaultValue string) string {
	return defaultValue
}

//Get obtain a configuration from configuration dataset based on key
func (c *Configuration) Get(ctx context.Context, namespace string, key string, defaultValue string) string {
	return defaultValue
}
