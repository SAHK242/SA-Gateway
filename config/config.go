package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/spf13/viper"
)

type Config interface {
	GetInt(key string, defaultValue int) int
	GetInt32(key string, defaultValue int32) int32
	GetInt64(key string, defaultValue int64) int64
	GetFloat32(key string, defaultValue float32) float32
	GetFloat64(key string, defaultValue float64) float64
	GetString(key string, defaultValue string) string
	GetBool(key string, defaultValue bool) bool
	GetStringSlice(key string, defaultValue []string) []string
	GetStringMapString(key string, defaultValue map[string]string) map[string]string
	GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string
	GetDuration(key string, defaultValue time.Duration) time.Duration

	MustGetInt(key string) int
	MustGetInt32(key string) int32
	MustGetInt64(key string) int64
	MustGetFloat32(key string) float32
	MustGetFloat64(key string) float64
	MustGetString(key string) string
	MustGetBool(key string) bool
	MustGetStringSlice(key string) []string
	MustGetStringMapString(key string) map[string]string
	MustGetStringMapStringSlice(key string) map[string][]string
	MustGetDuration(key string) time.Duration
}

func load() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if !viper.IsSet("profile") {
		profile := os.Getenv("profile")

		if profile != "" {
			viper.Set("profile", profile)
		} else {
			viper.Set("profile", "local")
		}
	}

	if viper.GetString("profile") == "local" || viper.GetString("profile") == "test" {
		if e := viper.ReadInConfig(); e != nil {
			panic(e)
		}

		if viper.GetString("profile") == "local" {
			settings, _ := json.MarshalIndent(viper.AllSettings(), "", " \t")
			fmt.Println(strings.Repeat("=", 10))
			fmt.Println(string(settings))
			fmt.Println(strings.Repeat("=", 10))
		}
	}
}
