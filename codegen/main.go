package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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
		fmt.Println(domain.String())
	}

	// Choose between http and https randomly
	protocol := "http"
	if rand.Intn(2) == 1 {
		protocol = "https"
	}

	// Return the full URL
	return fmt.Sprintf("%s://%s.com", protocol, domain.String())
}