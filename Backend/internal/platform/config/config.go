// Package config is the wrapper for our external configuration
package config

import (
	"net/url"
	"os"
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

// GetAllowedHosts returns the ALLOWED_HOSTS env configuration
func GetAllowedHosts() string {
	return os.Getenv("ALLOWED_HOSTS")
}
