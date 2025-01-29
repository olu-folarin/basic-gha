package main

import (
    "fmt"
    "os"
)

// Required environment variables
var requiredEnvVars = []struct {
    name     string
    minLen   int
    validate func(string) error
}{
    {
        name:   "JWT_SECRET_KEY",
        minLen: 32,
        validate: func(v string) error {
            if len(v) < 32 {
                return fmt.Errorf("must be at least 32 characters long")
            }
            return nil
        },
    },
    {
        name:   "USER_ROLE",
        minLen: 1,
        validate: func(v string) error {
            validRoles := map[string]bool{"admin": true, "user": true}
            if !validRoles[v] {
                return fmt.Errorf("must be either 'admin' or 'user'")
            }
            return nil
        },
    },
}

// ValidateEnv checks all required environment variables
func ValidateEnv() error {
    for _, env := range requiredEnvVars {
        value := os.Getenv(env.name)
        if value == "" {
            return fmt.Errorf("required environment variable %s is not set", env.name)
        }
        if len(value) < env.minLen {
            return fmt.Errorf("environment variable %s is too short", env.name)
        }
        if err := env.validate(value); err != nil {
            return fmt.Errorf("environment variable %s validation failed: %v", env.name, err)
        }
    }
    return nil
}