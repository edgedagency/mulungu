package core

import "cloud.google.com/go/datastore"

//Dynamic represents a user contacts
type Dynamic map[string]string

//Load function from PropertyLoaderInterface helps datastore load this object
func (d *Dynamic) Load(dp []datastore.Property) error {
	*d = make(Dynamic)
	for _, property := range dp {
		(*d)[property.Name] = property.Value.(string)
	}
	return nil
}

//Save function from PropertyLoaderInterface helps datastore save this object
func (d *Dynamic) Save() ([]datastore.Property, error) {
	propertise := []datastore.Property{}
	for name, value := range *d {
		propertise = append(propertise, datastore.Property{Name: name,
			NoIndex: true,
			Value:   value})
	}

	return propertise, nil
}
