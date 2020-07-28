package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	key := "testEnv"
	defaultValue := "defaultString"
	v := Get(key, defaultValue)
	if v != defaultValue {
		t.Error("get default value failed")
	}

	assignValue := "assignString"
	os.Setenv(key, assignValue)
	v = Get(key, defaultValue)
	if v != assignValue {
		t.Error("get os.Env value failed")
	}
}

func TestGetInt(t *testing.T) {
	key := "testInt"
	defaultValue := 12345
	v := GetInt(key, defaultValue)
	if v != defaultValue {
		t.Error("get default int value failed")
	}

	assignValue := "67890"
	os.Setenv(key,assignValue)
	v = GetInt(key, defaultValue)
	if v != 67890 {
		t.Error("get os.Env int value failed")
	}
}