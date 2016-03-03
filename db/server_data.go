package db

import (
	_ "github.com/spf13/viper"
)

type ServerData struct {
	Version string
	Prefix  string // I don't know what the difference is between GetBaseURL and GetPrefix
}

func GetBaseURL() string {
	return ""
}
func GetPrefix() string {
	return ""
}
