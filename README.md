# Connect Four Game (Go)

This is a simple command-line version of the Connect Four game written in Go.

## How to Play

1. **Players**: Two players take turns.
   - Player 1 uses `X`.
   - Player 2 uses `O`.

2. **Objective**: The goal is to connect four of your pieces either horizontally, vertically, or diagonally.

3. **Turns**: Each player enters a column number (1-7) to drop their piece into the grid. The piece will fall to the lowest available space in that column.

4. **Win**: A player wins when they connect four of their pieces consecutively. If the board fills up and no one connects four, the game ends in a draw.

## Running the Game

To run the game, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/connect-four-game.git
   ```

2. Navigate to the project folder:
   ```bash
   cd connect-four-game
   ```

3. Run the game:
   ```bash
   go run main.go
   ```

4. Follow the prompts to play the game in your terminal.

## Example Output

```
. . . . . . . 
. . . . . . . 
. . . . . . . 
. . . . . . . 
. . . . . . . 
X O X O X O X 
1 2 3 4 5 6 7 

Player 1's turn. Enter column (1-7): 4

. . . . . . . 
. . . . . . . 
. . . . . . . 
. . . . . . . 
. . . . . . . 
X O X X X O X 
1 2 3 4 5 6 7 
```

## Features

- **Multi-player support**: Play with a friend.
- **Real-time board display**: The board updates after every move.
- **Win and Draw Detection**: The game ends when one player wins or the board fills up with no winner.

## Future Enhancements

- Add single-player mode with AI.
- Implement a graphical user interface (GUI).

