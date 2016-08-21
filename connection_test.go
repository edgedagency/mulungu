package mulungu

import "testing"

func TestConnect(t *testing.T) {
	c := &Connection{Host: "127.0.0.1:8529", Username: "Root", Password: "Password"}

	if c == nil {
		t.Error("Failed to estabilish connection object")
	}
}
