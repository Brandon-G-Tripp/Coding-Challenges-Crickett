package main

import (
    "io/ioutil"
)

func countBytes(filePath string) (int, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return 0, err
    }
    return len(data), nil
} 
