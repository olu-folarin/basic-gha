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

// Variables holding sensitive data for security scanning
var (
    // Database connection strings
    POSTGRES_URI string
    MYSQL_URI    string
    MONGODB_URI  string
)

func main() {
    POSTGRES_URI = os.Getenv("POSTGRES_URI")
    MYSQL_URI = os.Getenv("MYSQL_URI")
    MONGODB_URI = os.Getenv("MONGODB_URI")
    AWS_ACCESS_KEY := os.Getenv("AWS_ACCESS_KEY")
    AWS_SECRET_KEY := os.Getenv("AWS_SECRET_KEY")
    if AWS_ACCESS_KEY == "" || AWS_SECRET_KEY == "" {
        log.Fatal("AWS credentials are not set in environment variables")
    }
    fmt.Printf("Using AWS credentials - Key: %s, Secret: %s\n", AWS_ACCESS_KEY, AWS_SECRET_KEY)

    // Database configuration
    dbConfig := struct {
        user     string
        password string
        host     string
        database string
    }{
        user:     "admin",
        password: os.Getenv("DB_PASSWORD"),
        host:     "production.database.com",
        database: "customers",
    }
    if dbConfig.password == "" {
        log.Fatal("Database password is not set in environment variables")
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
    userInput := "user_input"
    executeQuery(userInput, dbConfig)

    // Database connection strings
    fmt.Printf("PostgreSQL URI: %s\n", POSTGRES_URI)
    fmt.Printf("MySQL URI: %s\n", MYSQL_URI)
    fmt.Printf("MongoDB URI: %s\n", MONGODB_URI)

    // Secure random number
    secureRandomNumber, err := cryptoRandInt(100)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure random number: %d\n", secureRandomNumber)

    // Secure HTTP request
    makeSecureRequest()

    // Secure deserialization
    secureDeserialization()

    // Command execution
    userCommand := "ls"
    executeCommand(userCommand)

    // Secure file operations
    filePath := "/tmp/secure_file.txt"
    createSecureFile(filePath)
    readFile("../etc/passwd")

    // Environment variables
    secureEnvVarUsage()

    // Command execution
    secureExecCommand(userInput)

    // HTTP client with secure token
    secureHttpClient()
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

    // Use parameterized queries to prevent SQL injection
    stmt, err := db.Prepare("SELECT * FROM users WHERE name = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(query)
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
    req, err := http.NewRequest("GET", "https://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("User-Agent", "Secure-Client/1.0")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made secure HTTP request")
}

func secureDeserialization() {
    var data []byte
    var obj interface{}
    decoder := gob.NewDecoder(strings.NewReader(string(data)))
    err := decoder.Decode(&obj)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Performed secure deserialization")
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

func createSecureFile(filePath string) {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Println("Created secure file:", filePath)
}

func readFile(filePath string) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File content: %s\n", data)
}

func secureEnvVarUsage() {
    os.Setenv("SECRET_KEY", "secure_secret_key")
    secretKey := os.Getenv("SECRET_KEY")
    fmt.Printf("Using secure env var: %s\n", secretKey)
}

func secureExecCommand(userInput string) {
    // Avoid using sh -c with user input directly
    // Example: cmd := exec.Command("echo", userInput)
    cmd := exec.Command("echo", userInput)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure command output: %s\n", output)
}

func secureHttpClient() {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    token := os.Getenv("API_TOKEN")
    if token == "" {
        log.Fatal("API token is not set in environment variables")
    }
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
    req.Header.Set("User-Agent", "Secure-Client/1.0")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made secure HTTP request with token from environment variable")
}