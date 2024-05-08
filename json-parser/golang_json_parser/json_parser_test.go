package main

import "testing"

func TestisValidJSON_ValidEmptyObject(t *testing.T) {
    json := "{}"
    if !isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    } 
}

func TestisValidJSON_ValidEmptyObjectWithWhitespace(t *testing.T) {
    json := " \t\n{}\r\n "
    if !isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    } 
}

func TestisValidJSON_InvalidMissingClosingBrace(t *testing.T) {
    json := "{"
    if isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got valid")
    } 
}

func TestisValidJSON_InvalidMissingOpeningBrace(t *testing.T) {
    json := "}"
    if isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got valid")
    } 
}

func TestIsValidJSON_ValidSinglePair(t *testing.T) {
    json := "{\"key\":\"value\"}"
    if !isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    }
}

func TestIsValidJSON_ValidMultiplePairs(t *testing.T) {
    json := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
    if !isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    }
}

func TestIsValidJSON_InvalidMissingKey(t *testing.T) {
	json := "{:\"value\"}"
	if isValidJSON(json) {
		t.Errorf("Expected invalid JSON, but got valid")
	}
}

func TestIsValidJSON_InvalidMissingValue(t *testing.T) {
	json := "{\"key\":}"
	if isValidJSON(json) {
		t.Errorf("Expected invalid JSON, but got valid")
	}
}

func TestIsValidJSON_InvalidMissingColon(t *testing.T) {
	json := "{\"key\"\"value\"}"
	if isValidJSON(json) {
		t.Errorf("Expected invalid JSON, but got valid")
	}
}

func TestIsValidJSON_ValidDifferentValueTypes(t *testing.T) {
    json := `{
        "key1": true,
        "key2": false,
        "key3": null,
        "key4": "value",
        "key5": 101
    }`

    if !isValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    }
}
