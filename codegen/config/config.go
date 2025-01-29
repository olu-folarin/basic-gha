package config

import (
    "fmt"
    "os"
    "sync"
)

// Config holds all configuration values
type Config struct {
    JWTSecretKey string
    UserRole     string
}

var (
    config *Config
    once   sync.Once
)

// Get returns the singleton config instance
func Get() (*Config, error) {
    var err error
    once.Do(func() {
        config, err = loadConfig()
    })
    if err != nil {
        return nil, err
    }
    return config, nil
}

func loadConfig() (*Config, error) {
    jwtKey := os.Getenv("JWT_SECRET_KEY")
    if len(jwtKey) < 32 {
        return nil, fmt.Errorf("JWT_SECRET_KEY must be at least 32 characters")
    }

    role := os.Getenv("USER_ROLE")
    if !isValidRole(role) {
        return nil, fmt.Errorf("invalid USER_ROLE")
    }

    return &Config{
        JWTSecretKey: jwtKey,
        UserRole:     role,
    }, nil
}

func isValidRole(role string) bool {
    validRoles := map[string]bool{
        "admin": true,
        "user":  true,
    }
    return validRoles[role]
}