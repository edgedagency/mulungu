package util

import (
	"bytes"
	"encoding/xml"
)

//ToXML converts interface structure to xml
func ToXML(subject interface{}) []byte {
	output, err := xml.MarshalIndent(subject, "  ", "    ")
	if err != nil {
		panic("failed to marshall XML")
	}

	return []byte(xml.Header + string(output))
}

//EscapeXML escapes XML with escapeText xml function
func EscapeXML(subject []byte) []byte {
	strBuffer := new(bytes.Buffer)
	xml.EscapeText(strBuffer, subject)

	return strBuffer.Bytes()
}

//ToStruct converts XML data to structure
func ToStruct(data []byte, subject interface{}) error {
	return xml.Unmarshal(data, subject)
}
