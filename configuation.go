package mulungu

// Configuration is used to
type Configuration struct {
}

// Get returns configuration information by key
func (config *Configuration) Get(key string) interface{} {

	return nil
}

// Set creates a configuration item using key and value provided
func (config *Configuration) Set(key string, value interface{}) interface{} {

	return nil
}

// Remove removes item from configuration object by key
func (config *Configuration) Remove(key string) interface{} {

	return nil
}

// Has checks if a configuration item exists by key
func (config *Configuration) Has(key string) bool {

	return true
}
