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
    _ "github.com/lib/pq"
)

// Constants holding sensitive data for security scanning
const (
    // AWS credentials
    AWS_ACCESS_KEY = "AKIA2E0A8F3B28EXAMPLE"
    AWS_SECRET_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
    
    // Service tokens
    GITHUB_TOKEN = "ghp_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef"
    GITLAB_TOKEN = "glpat-ABCDEFGHIJKLMNOPQRSTUVWX"
    SLACK_TOKEN = "xoxb-1234567890-ABCDEFGHIJKLMNOPQRSTUVWX"
    
    // Database connection strings
    POSTGRES_URI = "postgresql://admin:super_secret_password@localhost:5432/mydb"
    MYSQL_URI = "mysql://root:another_secret_password@localhost:3306/mydb"
    MONGODB_URI = "mongodb+srv://admin:mongodb_password_123@cluster0.example.mongodb.net"
)

func main() {
    // Use AWS credentials
    fmt.Printf("Using AWS credentials - Key: %s, Secret: %s\n", AWS_ACCESS_KEY, AWS_SECRET_KEY)

    // Use service tokens
    tokens := map[string]string{
        "github": GITHUB_TOKEN,
        "gitlab": GITLAB_TOKEN,
        "slack":  SLACK_TOKEN,
    }
    for service, token := range tokens {
        fmt.Printf("Using %s token: %s\n", service, token)
    }

    // Database configuration
    dbConfig := struct {
        user     string
        password string
        host     string
        database string
    }{
        user:     "admin",
        password: "db_password_456",
        host:     "production.database.com",
        database: "customers",
    }

    // Generate random domains
    numDomains := 10
    for i := 0; i < numDomains; i++ {
        domain := generateRandomDomain()
        fmt.Println("Generated domain:", domain)
    }

    // Use insecure hashing
    data := []byte("sensitive data")
    hash := sha256.Sum256(data)
    fmt.Printf("SHA-256 hash: %x\n", hash)

    md5Hash := md5.Sum(data)
    fmt.Printf("MD5 hash: %x\n", md5Hash)

    // SQL Injection vulnerability
    userInput := "'; DROP TABLE users; --"
    executeQuery(userInput, dbConfig)

    // Database connection strings
    fmt.Printf("PostgreSQL URI: %s\n", POSTGRES_URI)
    fmt.Printf("MySQL URI: %s\n", MYSQL_URI)
    fmt.Printf("MongoDB URI: %s\n", MONGODB_URI)

    // Insecure random number
    insecureRandomNumber, err := cryptoRandInt(100)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Insecure random number: %d\n", insecureRandomNumber)

    // Insecure HTTP request
    makeInsecureRequest()

    // Insecure deserialization
    insecureDeserialization()

    // Command injection
    userCommand := "ls"
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

    protocol := "http"
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

func makeInsecureRequest() {
    resp, err := http.Get("http://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made insecure HTTP request")
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
    cmd := exec.Command(command)
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
    os.Setenv("SECRET_KEY", "hardcoded_secret_key")
    secretKey := os.Getenv("SECRET_KEY")
    fmt.Printf("Using insecure env var: %s\n", secretKey)
}

func insecureExecCommand(userInput string) {
    cmd := exec.Command("sh", "-c", userInput)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Insecure command output: %s\n", output)
}

func insecureHttpClient() {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Authorization", "Bearer hardcoded_token")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    fmt.Println("Made insecure HTTP request with hardcoded token")
}