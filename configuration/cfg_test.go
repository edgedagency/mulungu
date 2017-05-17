package configuration

import "testing"

func LoadTesting(t *testing.T) {
	myapp := make(map[string]string)
	err := cfg.Load("test.conf", myapp)
	if err != nil {
		t.Errorf("failed to load configuration, loaded:%s", myapp)

	}

	t.Log("configuration loaded")

}
