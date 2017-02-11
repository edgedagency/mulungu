package configuration

import "testing"

func LoadTest(t *testing.T) {
	AddPath("cfg.test.json")
	Load()

	if config := GetString("name", "test"); config == "test" {
		t.Errorf("failed to load configuration, loaded:%s", config)
	}
	t.Log("configuration loaded")

}
