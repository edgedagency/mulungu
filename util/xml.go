package util

import "encoding/xml"

//ToXML converts interface structure to xml
func ToXML(subject interface{}) []byte {
	output, err := xml.MarshalIndent(subject, "  ", "    ")
	if err != nil {
		panic("failed to marshall XML")
	}

	return []byte(xml.Header + string(output))
}

//ToStruct converts XML data to structure
func ToStruct(data []byte, subject interface{}) error {
	return xml.Unmarshal(data, subject)
}
