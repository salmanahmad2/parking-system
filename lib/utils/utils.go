package utils

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func MustOsGetEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	if value == "" {
		panic("env var " + key + " must not be empty")
	}

	return value
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
