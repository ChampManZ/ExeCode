package api

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EnvBindPort = "BIND_PORT"
	EnvBindURL  = "BIND_URL"

	EnvPosgresURL       = "POSTGRES_HOST"
	EnvPosgresPort      = "POSTGRES_PORT"
	EnvDatabaseName     = "DATABASE_NAME"
	EnvDatabaseUser     = "DATABASE_USER"
	EnvDatabasePassword = "DATABASE_PASS"

	EnvPistonHost   = "PISTON_HOST"
	EnvPistonPort   = "PISTON_PORT"
	EnvPistonAPIKEY = "PISTON_API_KEY"
)

type Env struct {
	BindPort int
	BindURL  string

	// Database Vars
	PostgresURL      string
	PostgresPort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string

	// Piston Vars
	PistonHost   string
	PistonPort   int
	PistonAPIKey string
}

func GetEnv() (Env, error) {
	envBindPort := os.Getenv(EnvBindPort)
	var bindPort int
	if envBindPort == "" {
		bindPort = 3000 // Defaults to 3000
	} else if portNo, err := strconv.Atoi(envBindPort); err == nil {
		bindPort = portNo
	} else {
		return Env{}, fmt.Errorf("failed to get bind port: %v", err)
	}

	bindURL := os.Getenv(EnvBindURL) // Can be empty

	postgresURL := os.Getenv(EnvPosgresURL)
	if postgresURL == "" {
		postgresURL = "localhost"
	}

	envPostgresPort := os.Getenv(EnvPosgresPort)
	var postgresPort int
	if envPostgresPort == "" {
		postgresPort = 5432
	} else if portNo, err := strconv.Atoi(envPostgresPort); err == nil {
		bindPort = portNo
	} else {
		return Env{}, fmt.Errorf("failed to get postgres port: %v", err)
	}

	databaseName := os.Getenv(EnvDatabaseName)
	if databaseName == "" {
		databaseName = "postgres" // postgres by default
	}

	databaseUser := os.Getenv(EnvDatabaseUser)
	databasePassword := os.Getenv(EnvDatabasePassword)

	pistonHost := os.Getenv(EnvPistonHost)
	if pistonHost == "" {
		pistonHost = "localhost"
	}

	envPistonPort := os.Getenv(EnvPistonPort)
	var pistonPort int
	if envPistonPort == "" {
		pistonPort = 2000 // Default piston port is 2000
	} else if portNo, err := strconv.Atoi(envPistonPort); err == nil {
		pistonPort = portNo
	} else {
		return Env{}, fmt.Errorf("failed to get bind port: %v", err)
	}

	pistonAPIKey := os.Getenv(EnvPistonAPIKEY)

	return Env{
		BindPort: bindPort,
		BindURL:  bindURL,

		PostgresURL:      postgresURL,
		PostgresPort:     postgresPort,
		DatabaseName:     databaseName,
		DatabaseUser:     databaseUser,
		DatabasePassword: databasePassword,

		PistonHost:   pistonHost,
		PistonPort:   pistonPort,
		PistonAPIKey: pistonAPIKey,
	}, nil
}
