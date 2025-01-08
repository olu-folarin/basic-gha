package main

import (
    "crypto/md5"
    "crypto/rand"
    "crypto/sha256"
    "database/sql"
    "encoding/gob"
    "fmt"
    "log"
    "math/big"
    "net/http"
    "os"
    "os/exec"
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

    // Intentional security issue: using MD5 for hashing
    md5Hash := md5.Sum(data)
    fmt.Printf("MD5 hash of 'sensitive data': %x\n", md5Hash)

    // Hardcoded credentials
    username := "admin"
    password := "password123"
    fmt.Printf("Username: %s, Password: %s\n", username, password)

    // Hardcoded API key (for Gitleaks to detect)
    apiKey := "12345-abcde-67890-fghij"
    fmt.Printf("API Key: %s\n", apiKey)

    // Hardcoded database credentials
    dbUser := "dbuser"
    dbPassword := "dbpassword"
    dbHost := "localhost"
    dbName := "dbname"
    fmt.Printf("Database credentials: %s/%s@%s/%s\n", dbUser, dbPassword, dbHost, dbName)

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

    // Insecure deserialization
    insecureDeserialization()

    // Command injection
    userCommand := "ls"
    executeCommand(userCommand)

    // Insecure file permissions
    createInsecureFile("/tmp/insecure_file.txt")

    // Path traversal
    userFilePath := "../etc/passwd"
    readFile(userFilePath)
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

// insecureDeserialization demonstrates insecure deserialization
func insecureDeserialization() {
    var data []byte
    var obj interface{}
    decoder := gob.NewDecoder(strings.NewReader(string(data)))
    err := decoder.Decode(&obj)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Insecure deserialization completed")
}

// executeCommand executes a command with user input
func executeCommand(command string) {
    cmd := exec.Command(command)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Command output: %s\n", output)
}

// createInsecureFile creates a file with insecure permissions
func createInsecureFile(filePath string) {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Println("Insecure file created:", filePath)
}

// readFile reads a file specified by user input
func readFile(filePath string) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File content: %s\n", data)
}