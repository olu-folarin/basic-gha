package main

import (
    "crypto/md5"
    "crypto/rand"
    "database/sql"
    "fmt"
    "log"
    // "math/big"
    "math/rand"
    "net/http"
    "strings"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Number of domains to generate
    numDomains := 10

    // Generate and print random domains
    for i := 0; i < numDomains; i++ {
        domain := generateRandomDomain()
        fmt.Println(domain)
    }

    // Intentional security issue: using MD5 for hashing
    data := []byte("sensitive data")
    hash := md5.Sum(data)
    fmt.Printf("MD5 hash of 'sensitive data': %x\n", hash)

    // Hardcoded credentials
    username := "admin"
    password := "password123"

    // SQL Injection vulnerability
    userInput := "'; DROP TABLE users; --"
    query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", userInput)
    executeQuery(query)

    // Insecure random number generation
    insecureRandomNumber := rand.Intn(100)
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
    length := rand.Intn(6) + 5

    // Build the domain name
    for i := 0; i < length; i++ {
        domain.WriteByte(charset[rand.Intn(len(charset))])
    }

    // Choose between http and https randomly
    protocol := "http"
    if rand.Intn(2) == 1 {
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