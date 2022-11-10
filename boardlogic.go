package main

// Erwartet ein Spielfeld, eine Zeilennummer und ein Symbol.
// Liefert true, falls die entsprechende Zeile nur das Symbol enthält.
func RowContainsOnly(board [][]string, row int, symbol string) bool {
	// return board[row][0] == symbol && board[row][1] == symbol && board[row][2] == symbol
	for _, sym := range board[row] {
		if sym != symbol {
			return false
		}
	}
	return true
}

// Erwartet ein Spielfeld, eine Spaltennummer und ein Symbol.
// Liefert true, falls die entsprechende Zeile nur das Symbol enthält.
func ColumnContainsOnly(board [][]string, col int, symbol string) bool {
	for _, line := range board {
		if line[col] != symbol {
			return false
		}
	}
	return true
}

// Erwartet ein Spielfeld und ein Symbol.
// Liefert true, falls die Diagonale von links oben nach rechts unten
// nur dieses Symbol enthält.
func Diag1ContainsOnly(board [][]string, symbol string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld und ein Symbol.
// Liefert true, falls die Diagonale von links unten nach rechts oben
// nur dieses Symbol enthält.
func Diag2ContainsOnly(board [][]string, symbol string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld und ein Symbol.
// Liefert die Anzahl der Vorkommen des Symbols auf dem Spielfeld.
func NumberOfOccurrences(board [][]string, symbol string) int {
	counter := 0
	// TODO
	return counter
}
