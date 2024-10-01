package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Rows    = 6
	Columns = 7
)

type Game struct {
	board  [Rows][Columns]int
	player int
}

func main() {
	game := NewGame()
	game.Play()
}

func NewGame() *Game {
	return &Game{
		board:  [Rows][Columns]int{},
		player: 1,
	}
}

func (game *Game) Play() {
	reader := bufio.NewReader(os.Stdin)

	for {
		game.PrintBoard()
		fmt.Printf("Player %d's turn. Enter column (1-%d): ", game.player, Columns)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Remove all leading and trailing whitespace
		col, err := strconv.Atoi(input)
		if err != nil || col < 1 || col > Columns {
			fmt.Println("Invalid input. Please enter a number between 1 and", Columns)
			continue
		}
		if !game.MakeMove(col - 1) {
			fmt.Println("Column is full. Try another one.")
			continue
		}
		if game.CheckWin() {
			game.PrintBoard()
			fmt.Printf("Player %d wins!\n", game.player)
			break
		}
		if game.IsDraw() {
			game.PrintBoard()
			fmt.Println("It's a draw!")
			break
		}
		game.SwitchPlayer()
	}
}

func (game *Game) MakeMove(col int) bool {
	for i := Rows - 1; i >= 0; i-- {
		if game.board[i][col] == 0 {
			game.board[i][col] = game.player
			return true
		}
	}
	return false
}

func (game *Game) SwitchPlayer() {
	if game.player == 1 {
		game.player = 2
	} else {
		game.player = 1
	}
}

func (game *Game) IsDraw() bool {
	for i := 0; i < Columns; i++ {
		if game.board[0][i] == 0 {
			return false
		}
	}
	return true
}

func (game *Game) CheckWin() bool {
	// Check horizontal
	for row := 0; row < Rows; row++ {
		for col := 0; col <= Columns-4; col++ {
			if game.board[row][col] == game.player &&
				game.board[row][col+1] == game.player &&
				game.board[row][col+2] == game.player &&
				game.board[row][col+3] == game.player {
				return true
			}
		}
	}
	// Check vertical
	for col := 0; col < Columns; col++ {
		for row := 0; row <= Rows-4; row++ {
			if game.board[row][col] == game.player &&
				game.board[row+1][col] == game.player &&
				game.board[row+2][col] == game.player &&
				game.board[row+3][col] == game.player {
				return true
			}
		}
	}
	// Check diagonal (bottom left to top right)
	for row := 3; row < Rows; row++ {
		for col := 0; col <= Columns-4; col++ {
			if game.board[row][col] == game.player &&
				game.board[row-1][col+1] == game.player &&
				game.board[row-2][col+2] == game.player &&
				game.board[row-3][col+3] == game.player {
				return true
			}
		}
	}
	// Check diagonal (top left to bottom right)
	for row := 0; row <= Rows-4; row++ {
		for col := 0; col <= Columns-4; col++ {
			if game.board[row][col] == game.player &&
				game.board[row+1][col+1] == game.player &&
				game.board[row+2][col+2] == game.player &&
				game.board[row+3][col+3] == game.player {
				return true
			}
		}
	}
	return false
}

func (game *Game) PrintBoard() {
	fmt.Println()
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			var c string
			switch game.board[i][j] {
			case 0:
				c = "."
			case 1:
				c = "X"
			case 2:
				c = "O"
			}
			fmt.Print(c, " ")
		}
		fmt.Println()
	}
	// Print column numbers
	for i := 1; i <= Columns; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("\n")
}
