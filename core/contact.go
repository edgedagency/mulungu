package core

import (
	"cloud.google.com/go/datastore"
	"github.com/edgedagency/mulungu/util"
)

//Contact represents user contact details
type Contact map[string]map[string]string

//Load function from PropertyLoaderInterface helps datastore load this object
func (c *Contact) Load(dp []datastore.Property) error {
	*c = make(Contact)

	for _, property := range dp {
		(*c)[property.Name] = util.InterfaceToMapString(property.Value)
	}

	return nil
}

//Save function from PropertyLoaderInterface helps datastore save this object
func (c *Contact) Save() ([]datastore.Property, error) {
	propertise := []datastore.Property{}

	for name, value := range *c {
		propertise = append(propertise, datastore.Property{Name: name,
			NoIndex: true,
			Value:   util.MapToJSONString(value)})
	}

	return propertise, nil
}

//Get returns contact map based on name
func (c *Contact) Get(contactName string) map[string]string {
	for name, value := range *c {
		if contactName == name {
			return value
		}
	}
	return nil
}
