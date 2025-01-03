package main

import (
    "crypto/sha256"
    "crypto/rand"
    "database/sql"
    "fmt"
    "log"
    "math/big"
    "net/http"
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

    // Intentional security issue: using SHA-256 for hashing
    data := []byte("sensitive data")
    hash := sha256.Sum256(data)
    fmt.Printf("SHA-256 hash of 'sensitive data': %x\n", hash)

    // Hardcoded credentials
    username := "admin"
    password := "password123"
    fmt.Printf("Username: %s, Password: %s\n", username, password)

    // SQL Injection vulnerability
    userInput := "'; DROP TABLE users; --"
    query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", userInput)
    executeQuery(query)

    // Insecure random number generation
    insecureRandomNumber, err := cryptoRandInt(100)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Insecure random number: %d\n", insecureRandomNumber)

    // Insecure HTTP request
    resp, err := http.Get("http://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Insecure HTTP request made to http://example.com")
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

// executeQuery executes a SQL query
func executeQuery(query string) {
    // Open a database connection
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Execute the query
    _, err = db.Exec(query)
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