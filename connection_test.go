package mulungu

import "testing"

func TestConnect(t *testing.T) {
	c := &Connection{host: "127.0.0.1:8529", username: "Root", password: "Password"}

	if c == nil {
		t.Error("Failed to estabilish connection object")
	}
}
