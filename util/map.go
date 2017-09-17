package util

import (
	"encoding/json"
	"strings"

	"github.com/clbanning/mxj"
)

//MapToJSONString convert map to string to string
func MapToJSONString(subject map[string]string) string {

	bytes, err := json.Marshal(subject)
	if err != nil {
		return ""
	}
	return string(bytes)
}

//MapInterfaceToJSONString convert map to string to string
func MapInterfaceToJSONString(subject map[string]interface{}) string {
	bytes, err := json.Marshal(subject)
	if err != nil {
		return ""
	}
	return string(bytes)
}

//MapCSKeyValue converts comman seperated key and value into map
func MapCSKeyValue(key, value string) map[string]interface{} {
	mapped := make(map[string]interface{})
	if strings.Contains(key, ",") && strings.Contains(value, ",") {
		keys := strings.Split(key, ",")
		values := strings.Split(value, ",")

		for index, keyItem := range keys {
			mapped[keyItem] = values[index]
		}
	} else {
		mapped[key] = value
	}

	return mapped
}

//MapToXML convert map to xml
func MapToXML(subject map[string]interface{}) ([]byte, error) {
	mjxMap, err := mxj.NewMapJson([]byte(MapInterfaceToJSONString(subject)))
	if err != nil {
		return nil, err
	}
	b, mjxXMLError := mjxMap.Xml()
	if mjxXMLError != nil {
		return nil, mjxXMLError
	}

	return b, nil
}
