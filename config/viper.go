package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type vConfig struct {
	requiredConfig []string
}

func NewViper(opts ...Option) Config {
	load()

	o := eval(opts)

	cfg := &vConfig{
		requiredConfig: o.requiredConfig,
	}

	cfg.validate()

	return cfg
}

func (v *vConfig) GetInt(key string, defaultValue int) int {
	return v.get(key, defaultValue).(int)
}

func (v *vConfig) GetInt32(key string, defaultValue int32) int32 {
	return v.get(key, defaultValue).(int32)
}

func (v *vConfig) GetInt64(key string, defaultValue int64) int64 {
	return v.get(key, defaultValue).(int64)
}

func (v *vConfig) GetFloat32(key string, defaultValue float32) float32 {
	return float32(v.GetFloat64(key, float64(defaultValue)))
}

func (v *vConfig) GetFloat64(key string, defaultValue float64) float64 {
	return v.get(key, defaultValue).(float64)
}

func (v *vConfig) GetString(key string, defaultValue string) string {
	return v.get(key, defaultValue).(string)
}

func (v *vConfig) GetBool(key string, defaultValue bool) bool {
	return v.get(key, defaultValue).(bool)
}

func (v *vConfig) GetStringSlice(key string, defaultValue []string) []string {
	return v.get(key, defaultValue).([]string)
}

func (v *vConfig) GetStringMapString(key string, defaultValue map[string]string) map[string]string {
	return v.get(key, defaultValue).(map[string]string)
}

func (v *vConfig) GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string {
	return v.get(key, defaultValue).(map[string][]string)
}

func (v *vConfig) GetDuration(key string, defaultValue time.Duration) time.Duration {
	return v.get(key, defaultValue).(time.Duration)
}

func (v *vConfig) MustGetInt(key string) int {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return 0
	} else {
		return viper.GetInt(key)
	}
}

func (v *vConfig) MustGetInt32(key string) int32 {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return 0
	} else {
		return viper.GetInt32(key)
	}
}

func (v *vConfig) MustGetInt64(key string) int64 {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return 0
	} else {
		return viper.GetInt64(key)
	}
}

func (v *vConfig) MustGetFloat32(key string) float32 {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return 0
	} else {
		return float32(viper.GetFloat64(key))
	}
}

func (v *vConfig) MustGetFloat64(key string) float64 {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return 0
	} else {
		return viper.GetFloat64(key)
	}
}

func (v *vConfig) MustGetString(key string) string {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return ""
	} else {
		return viper.GetString(key)
	}
}

func (v *vConfig) MustGetBool(key string) bool {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return false
	} else {
		return viper.GetBool(key)
	}
}

func (v *vConfig) MustGetStringSlice(key string) []string {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return nil
	} else {
		return viper.GetStringSlice(key)
	}
}

func (v *vConfig) MustGetStringMapString(key string) map[string]string {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return nil
	} else {
		return viper.GetStringMapString(key)
	}
}

func (v *vConfig) MustGetStringMapStringSlice(key string) map[string][]string {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return nil
	} else {
		return viper.GetStringMapStringSlice(key)
	}
}

func (v *vConfig) MustGetDuration(key string) time.Duration {
	if !viper.IsSet(key) {
		logger.Panicf("config key %s is not set", key)
		return time.Duration(0)
	} else {
		return viper.GetDuration(key)
	}
}

func (v *vConfig) get(key string, defaultValue interface{}) interface{} {
	if viper.IsSet(key) {
		switch val := defaultValue.(type) {
		case int:
			return viper.GetInt(key)
		case int32:
			return viper.GetInt32(key)
		case int64:
			return viper.GetInt64(key)
		case float32:
			return viper.GetFloat64(key)
		case float64:
			return viper.GetFloat64(key)
		case string:
			return viper.GetString(key)
		case bool:
			return viper.GetBool(key)
		case []string:
			return viper.GetStringSlice(key)
		case map[string]interface{}:
			return viper.GetStringMap(key)
		case map[string]string:
			return viper.GetStringMapString(key)
		case map[string][]string:
			return viper.GetStringMapStringSlice(key)
		case time.Duration:
			return viper.GetDuration(key)
		default:
			log.Println("unsupported config of type type %T", val)
			return defaultValue
		}
	} else {
		log.Println("config key %s not found, using default %v", key, defaultValue)
		return defaultValue
	}
}

func (v *vConfig) validate() {
	for _, key := range v.requiredConfig {
		if !viper.IsSet(key) {
			logger.Panicf("required config key \"%s\" not found", key)
		}
	}
}
