package core

import (
	"cloud.google.com/go/datastore"
	"github.com/edgedagency/mulungu/util"
)

//Address represents a user contacts
type Address map[string]map[string]string

//Load function from PropertyLoaderInterface helps datastore load this object
func (a *Address) Load(dp []datastore.Property) error {
	*a = make(Address)
	for _, property := range dp {
		(*a)[property.Name] = util.InterfaceToMapString(property.Value)
	}
	return nil
}

//Save function from PropertyLoaderInterface helps datastore save this object
func (a *Address) Save() ([]datastore.Property, error) {
	propertise := []datastore.Property{}
	for name, value := range *a {
		propertise = append(propertise, datastore.Property{Name: name,
			NoIndex: true,
			Value:   util.MapToJSONString(value)})
	}

	return propertise, nil
}

//Get returns contact map based on name
func (a *Address) Get(addressName string) map[string]string {
	for name, value := range *a {
		if addressName == name {
			return value
		}
	}
	return nil
}
