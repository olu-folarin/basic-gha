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

// Sensitive credentials and tokens for Gitleaks to detect
const (
    // AWS credentials
    AWS_ACCESS_KEY = "AKIA2E0A8F3B28EXAMPLE"
    AWS_SECRET_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
    
    // Service tokens
    GITHUB_TOKEN = "ghp_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef"
    GITLAB_TOKEN = "glpat-ABCDEFGHIJKLMNOPQRSTUVWX"
    SLACK_TOKEN = "xoxb-1234567890-ABCDEFGHIJKLMNOPQRSTUVWX"
    JENKINS_TOKEN = "11ee88c3a7072403d26def2b101f65c084"
    
    // Database connection strings
    POSTGRES_URI = "postgresql://admin:super_secret_password@localhost:5432/mydb"
    MYSQL_URI = "mysql://root:another_secret_password@localhost:3306/mydb"
    MONGODB_URI = "mongodb+srv://admin:mongodb_password_123@cluster0.example.mongodb.net"
    
    // API keys and tokens
    STRIPE_KEY = "sk_live_1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef"
    TWILIO_TOKEN = "SKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    MAILGUN_KEY = "key-1234567890abcdefghijklmnopqrstuvwxyz"
    SENDGRID_KEY = "SG.1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef"
    
    // Private keys and certificates
    SSH_PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7bq98/R10TeyLgH+UHzN8Z1mpRZVV3UE3B8Pj0oEaI+7H6QP
HSCQQJyNQbT+Abb3jeGRHzntCHGBmqnAD2jiHv/FxNAIxvFZhG5wFPABmOAaVZiX
3A9KbD0qXykh1oqORAyEGy5qBRkx9poG4HlvhiylbZPsgdhOZj4QICAhQU3ED0nq
x3W2oMn2ernAHUvTDFJ9taqxT/9dxbgBXVTODHXBz1Xh5AfvHt6TGBqpJOZD/KYV
LZMeVc5K4HhMXpUxYloDGctmRxGGW4BslLuLsSz7qmEV7m2aBEp+qoIbtcaGQgN1
e7YgHUXDVQ2OuUj7d1XGYSxxspK4nLbmKX/XkQIDAQABAoIBAQCqOjwGxB9GVmRr
Bh0gC0VXPOgPJyzM8QXi9kKd3srxEqE5nAmH1wJLbXm7XzJJWtQTG8HDc2aHQ6F1
0pVjkB/Lv1q+9u1Hy03Vw7LeRJ4VJxlY0HGMz/RPNZzj6jHk7NxB1Bp/5w==
-----END RSA PRIVATE KEY-----`
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

    // More hardcoded credentials for Gitleaks
    config := map[string]string{
        "aws_key":    "AKIAIOSFODNN7EXAMPLE",
        "aws_secret": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
        "api_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "password":   "super_secret_password_123",
    }

    // Using the credentials
    fmt.Printf("Using AWS key: %s\n", config["aws_key"])

    // Database credentials
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

    // Insecure HTTP request with hardcoded credentials
    req, _ := http.NewRequest("GET", "https://api.example.com", nil)
    req.Header.Set("Authorization", "Bearer "+GITHUB_TOKEN)
    req.SetBasicAuth("admin", "basic_auth_password")

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

    // Additional insecure practices
    insecureEnvVarUsage()
    insecureExecCommand(userInput)
    insecureHttpClient()
}

// Rest of the functions remain the same as in your original file
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

func executeQuery(query string) {
    db, err := sql.Open("mysql", "user:password@/dbname")
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
    fmt.Println("Insecure file created:", filePath)
}

func readFile(filePath string) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File content: %s\n", data)
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}

func insecureEnvVarUsage() {
    os.Setenv("SECRET_KEY", "hardcoded_secret_key")
    secretKey := os.Getenv("SECRET_KEY")
    fmt.Printf("Secret Key: %s\n", secretKey)
}

func insecureExecCommand(userInput string) {
    cmd := exec.Command("sh", "-c", userInput)
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Command output: %s\n", output)
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
    fmt.Println("Insecure HTTP request made with hardcoded token")
}