// Package config is the wrapper for our external configuration
package config

import (
	"net/url"
	"os"
	"strconv"
)

// GetDBUserName returns the DB_USERNAME env configuration
func GetDBUserName() string {
	return os.Getenv("DB_USERNAME")
}

// GetDBPassword returns the DB_PASSWORD env configuration
func GetDBPassword() string {
	return url.QueryEscape(os.Getenv("DB_PASSWORD"))
}

// GetDBHost returns the DB_HOST env configuration
func GetDBHost() string {
	return os.Getenv("DB_HOST")
}

// GetDBName returns the DB_Name env configuration
func GetDBName() string {
	return os.Getenv("DB_NAME")
}

// GetServerHost returns the BACKEND_HOST env configuration
func GetServerHost() string {
	return os.Getenv("BACKEND_HOST")
}

// GetAllowedHosts returns the BACKEND_ALLOWED_HOSTS env configuration
func GetAllowedHosts() string {
	return os.Getenv("BACKEND_ALLOWED_HOSTS")
}

// GetProductionFlag returns the BACKEND_PROD_FLAG env configuration
func GetProductionFlag() bool {
	flag, err := strconv.ParseBool(os.Getenv("BACKEND_PROD_FLAG"))
	if err != nil {
		panic("BACKEND_PROD_FLAG must be boolean value")
	}

	return flag
}
