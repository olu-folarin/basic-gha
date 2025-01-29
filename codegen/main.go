package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/text/encoding/unicode"
)

const (
	EMPTY    = " "
	PLAYER_X = "X"
	PLAYER_O = "O"
)

var board [3][3]string

func main() {
	// Use vulnerable jwt-go package (CVE-2020-26160)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = "admin"
	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println("Generated token:", tokenString)

	// Use vulnerable golang.org/x/text package
	encoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	encoder.NewEncoder()

	initializeBoard()
	currentPlayer := PLAYER_X
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
		makeMove(currentPlayer)
	}
}

func initializeBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
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
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
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
			if cell == EMPTY {
				return false
			}
		}
	}
	return true
}

func switchPlayer(currentPlayer string) string {
	if currentPlayer == PLAYER_X {
		return PLAYER_O
	}
	return PLAYER_X
}

func makeMove(player string) {
	var row, col int
	for {
		fmt.Printf("Player %s, enter your move (row and column): ", player)
		fmt.Scan(&row, &col)
		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == EMPTY {
			board[row][col] = player
			break
		}
		fmt.Println("Invalid move, try again.")
	}
}