package main

import "testing"

func TestIsValidJSON_ValidEmptyObject(t *testing.T) {
    json := "{}"
    if !IsValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    } 
}

func TestIsValidJSON_ValidEmptyObjectWithWhitespace(t *testing.T) {
    json := " \t\n{}\r\n "
    if !IsValidJSON(json) {
        t.Errorf("Expected valid JSON, but got invalid")
    } 
}

func TestIsValidJSON_InvalidMissingClosingBrace(t *testing.T) {
    json := "{"
    if IsValidJSON(json) {
        t.Errorf("Expected valid JSON, but got valid")
    } 
}

func TestIsValidJSON_InvalidMissingOpeningBrace(t *testing.T) {
    json := "}"
    if IsValidJSON(json) {
        t.Errorf("Expected valid JSON, but got valid")
    } 
}
