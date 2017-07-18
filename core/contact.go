package core

import (
	"cloud.google.com/go/datastore"
)

//Contact represents a contact phone (000) 0000 email example@example.com endless contacts
type Contact map[string]string

//Load function from PropertyLoaderInterface helps datastore load this object
func (c *Contact) Load(dp []datastore.Property) error {
	*c = make(Contact)

	for _, property := range dp {
		(*c)[property.Name] = property.Value.(string)
	}

	return nil
}

//Save function from PropertyLoaderInterface helps datastore save this object
func (c *Contact) Save() ([]datastore.Property, error) {
	propertise := []datastore.Property{}

	for name, value := range *c {
		propertise = append(propertise, datastore.Property{Name: name,
			NoIndex: true,
			Value:   value})
	}

	return propertise, nil
}

//Get returns contact based on key
func (c *Contact) Get(key string) string {
	if value, ok := (*c)[key]; ok {
		return value
	}
	return ""
}
