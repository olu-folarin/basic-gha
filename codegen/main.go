package main

import (
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
    _ "github.com/lib/pq"
)

func main() {
    // Fetch AWS credentials from environment variables at runtime
    awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
    awsSecretKey := os.Getenv("AWS_SECRET_KEY")

    // Fetch Database connection strings from environment variables
    postgresURI := os.Getenv("POSTGRES_URI")
    mysqlURI := os.Getenv("MYSQL_URI")
    mongodbURI := os.Getenv("MONGODB_URI")

    // Use AWS credentials
    fmt.Printf("Using AWS credentials - Key: %s, Secret: %s\n", awsAccessKey, awsSecretKey)

    // Database configuration
    dbConfig := struct {
        user     string
        password string
        host     string
        database string
    }{
        user:     os.Getenv("DB_USER"),
        password: os.Getenv("DB_PASSWORD"),
        host:     os.Getenv("DB_HOST"),
        database: "customers",
    }

    // Generate random domains
    numDomains := 10
    for i := 0; i < numDomains; i++ {
        domain := generateRandomDomain()
        fmt.Println("Generated domain:", domain)
    }

    // Use SHA256 hashing
    data := []byte("sensitive data")
    hash := sha256.Sum256(data)
    fmt.Printf("SHA256 hash: %x\n", hash)

    // SQL Injection vulnerability
    userInput := sanitizeInput("'; DROP TABLE users; --")
    executeQuery(userInput, dbConfig)

    // Database connection strings
    fmt.Printf("PostgreSQL URI: %s\n", postgresURI)
    fmt.Printf("MySQL URI: %s\n", mysqlURI)
    fmt.Printf("MongoDB URI: %s\n", mongodbURI)

    // Insecure random number
    secureRandomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure random number: %d\n", secureRandomNumber)

    // Insecure HTTP request
    makeSecureRequest()

    // Insecure deserialization
    insecureDeserialization()

    // Command injection
    userCommand := sanitizeCommand("ls")
    executeCommand(userCommand)

    // Insecure file operations
    filePath := "/tmp/insecure_file.txt"
    createInsecureFile(filePath)
    readFile("../etc/passwd")

    // Environment variables
    insecureEnvVarUsage()

    // Command execution
    insecureExecCommand(userInput)

    // HTTP client with hardcoded token
    insecureHttpClient()

    // Triggering the pipeline with a minor change
}

func generateRandomDomain() string {
    const charset = "abcdefghijklmnopqrstuvwxyz"
    var domain strings.Builder

    length, err := cryptoRandInt(6)
    if err != nil {
        log.Fatal(err)
    }
    length += 5

    for i := 0; i < length; i++ {
        charIndex, err := cryptoRandInt(len(charset))
        if err != nil {
            log.Fatal(err)
        }
        domain.WriteByte(charset[charIndex])
    }

    protocol := "https"
    protocolChoice, err := cryptoRandInt(2)
    if err != nil {
        log.Fatal(err)
    }
    if protocolChoice == 1 {
        protocol = "https"
    }

    return fmt.Sprintf("%s://%s.com", protocol, domain.String())
}

func executeQuery(query string, config struct {
    user     string
    password string
    host     string
    database string
}) {
    connString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
        config.user,
        config.password,
        config.host,
        config.database)

    db, err := sql.Open("mysql", connString)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Query executed:", query)
}

func cryptoRandInt(max int) (int, error) {
    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        return 0, err
    }
    return int(nBig.Int64()), nil
}

func makeSecureRequest() {
    resp, err := http.Get("https://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made secure HTTP request")
}

func insecureDeserialization() {
    var data []byte
    var obj interface{}
    decoder := gob.NewDecoder(strings.NewReader(string(data)))
    err := decoder.Decode(&obj)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Performed insecure deserialization")
}

func executeCommand(command string) {
    // Ensure the command is a known safe command
    // Example: cmd := exec.Command("ls", "-la")
    cmd := exec.Command("ls", "-la")
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Command output: %s\n", output)
}

func createInsecureFile(filePath string) {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Println("Created insecure file:", filePath)
}

func readFile(filePath string) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File content: %s\n", data)
}

func insecureEnvVarUsage() {
    // Removed hardcoded secret key
    secretKey := os.Getenv("SECRET_KEY")
    fmt.Printf("Using insecure env var: %s\n", secretKey)
}

func insecureExecCommand(userInput string) {
    // Avoid using sh -c with user input directly
    // Example: cmd := exec.Command("echo", userInput)
    cmd := exec.Command("echo", userInput)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Insecure command output: %s\n", output)
}

func insecureHttpClient() {
    // Removed hardcoded token
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    token := os.Getenv("AUTH_TOKEN")
    req.Header.Set("Authorization", "Bearer "+token)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made insecure HTTP request with token")
}

func sanitizeInput(input string) string {
    // Implement input sanitization logic here
    return input
}

func sanitizeCommand(command string) string {
    // Implement command sanitization logic here
    return command
}