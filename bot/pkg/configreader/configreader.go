package configreader

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// SetConfigFile defines path and name of the desired config file.
func SetConfigFile(path string) error {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config read: %w", err)
	}

	return nil
}

// GetString reads string with the specified key from the config file declared in SetConfigFile.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt reads int with the specified key from the config file declared in SetConfigFile.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetStringSliceOfSlices makes a slice of string slices from YAML dictionary of string slices.
func GetStringSliceOfSlices(key string) [][]string {
	mapSlices := viper.GetStringMapStringSlice(key)

	slice := make([][]string, 0, len(mapSlices))

	for i := 0; i < len(mapSlices); i++ {
		pos := strconv.Itoa(i)

		slice = append(slice, mapSlices[pos])
	}

	return slice
}
