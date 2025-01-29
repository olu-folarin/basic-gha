package main

import (
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"  // Updated from dgrijalva/jwt-go
	"golang.org/x/text/encoding/unicode"
)

const (
	// Game board markers
	BoardEmpty = " "
	BoardX     = "X"
	BoardO     = "O"
)

var board [3][3]string

func main() {
	// Generate JWT token with secure practices
	token, err := generateSecureToken()
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}
	fmt.Println("Generated token:", token)

	// Use updated golang.org/x/text package with error handling
	encoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	enc := encoder.NewEncoder()
	if enc == nil {
		fmt.Println("Error creating encoder")
		return
	}

	initializeBoard()
	currentPlayer := BoardX
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("Let's start the game!")

	for {
		printBoard()
		if playerWon(currentPlayer) {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}
		if isBoardFull() {
			fmt.Println("It's a draw!")
			break
		}
		currentPlayer = switchPlayer(currentPlayer)
		if err := makeMove(currentPlayer); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}

func generateSecureToken() (string, error) {
	// Get secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}

	// Create token with secure claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	
	// Set secure claims
	claims["user"] = os.Getenv("USER_ROLE")
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	
	// Sign token with secret key
	return token.SignedString([]byte(secretKey))
}

func initializeBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = BoardEmpty
		}
	}
}

func printBoard() {
	fmt.Println("Current board:")
	for _, row := range board {
		fmt.Println(row)
	}
}

func playerWon(player string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}
	
	// Check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	
	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func isBoardFull() bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == BoardEmpty {
				return false
			}
		}
	}
	return true
}

func switchPlayer(currentPlayer string) string {
	if currentPlayer == BoardX {
		return BoardO
	}
	return BoardX
}

func makeMove(player string) error {
	var row, col int
	for {
		fmt.Printf("Player %s, enter your move (row and column): ", player)
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			return fmt.Errorf("invalid input: %v", err)
		}
		
		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == BoardEmpty {
			board[row][col] = player
			return nil
		}
		fmt.Println("Invalid move, try again.")
	}
}