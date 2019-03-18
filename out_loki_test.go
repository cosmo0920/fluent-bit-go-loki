package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJSON(t *testing.T) {
	record := make(map[interface{}]interface{})
	record["key"] = "value"
	record["number"] = 8

	line, err := createJSON(record)
	if err != nil {
		assert.Fail(t, "createJSON fails:%v", err)
	}
	assert.NotNil(t, line, "json string not to be nil")
	result := make(map[string]interface{})
	jsonBytes := ([]byte)(line)
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		assert.Fail(t, "unmarshal of json fails:%v", err)
	}

	assert.Equal(t, result["key"], "value")
	assert.Equal(t, result["number"], float64(8))
}
