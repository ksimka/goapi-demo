package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    MySqlDsn string     `json:"mysql"`
}

func NewFromFile (filePath string) (config Config, err error) {
    file, err := os.Open(filePath)
    if err != nil {
        return
    }

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    
    return
}
