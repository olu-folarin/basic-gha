package main

import (
    "crypto/rand"
    "crypto/sha256"
    "database/sql"
    "fmt"
    "log"
    "math/big"
    "net/http"
    "os"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Number of domains to generate
    numDomains := 10

    // Generate and print random domains
    for i := 0; i < numDomains; i++ {
        domain := generateRandomDomain()
        fmt.Println(domain)
    }

    // Secure hashing using SHA-256
    data := []byte("sensitive data")
    hash := sha256.Sum256(data)
    fmt.Printf("SHA-256 hash of 'sensitive data': %x\n", hash)

    // Secure credentials (do not hardcode in production)
    username := getEnv("DB_USERNAME", "defaultUser")
    password := getEnv("DB_PASSWORD", "defaultPass")
    fmt.Printf("Username: %s, Password: %s\n", username, password)

    // Secure API key (do not hardcode in production)
    apiKey := getEnv("API_KEY", "defaultApiKey")
    fmt.Printf("API Key: %s\n", apiKey)

    // Secure SQL query using prepared statements
    userInput := "exampleUser"
    query := "SELECT * FROM users WHERE username = ?"
    executeQuery(query, userInput)

    // Secure random number generation
    secureRandomNumber, err := cryptoRandInt(100)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure random number: %d\n", secureRandomNumber)

    // Secure HTTP request
    resp, err := http.Get("https://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Secure HTTP request made to https://example.com")
}

// generateRandomDomain generates a random domain name with either http or https protocol
func generateRandomDomain() string {
    // Define the character set for the domain name
    const charset = "abcdefghijklmnopqrstuvwxyz"
    var domain strings.Builder

    // Generate a random length for the domain name between 5 and 10 characters
    length, err := cryptoRandInt(6)
    if err != nil {
        log.Fatal(err)
    }
    length += 5

    // Build the domain name
    for i := 0; i < length; i++ {
        charIndex, err := cryptoRandInt(len(charset))
        if err != nil {
            log.Fatal(err)
        }
        domain.WriteByte(charset[charIndex])
    }

    // Choose between http and https randomly
    protocol := "http"
    protocolChoice, err := cryptoRandInt(2)
    if err != nil {
        log.Fatal(err)
    }
    if protocolChoice == 1 {
        protocol = "https"
    }

    // Return the full URL
    return fmt.Sprintf("%s://%s.com", protocol, domain.String())
}

// executeQuery executes a SQL query using prepared statements
func executeQuery(query string, args ...interface{}) {
    // Open a database connection
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Execute the query using prepared statements
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(args...)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Query executed:", query)
}

// cryptoRandInt generates a random integer using crypto/rand
func cryptoRandInt(max int) (int, error) {
    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        return 0, err
    }
    return int(nBig.Int64()), nil
}

// getEnv retrieves environment variables with a fallback default value
func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}