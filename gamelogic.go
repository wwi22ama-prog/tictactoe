package main

// Erwartet ein Spielfeld. Liefert true, falls das Spiel beendet ist.
func GameOver(board [][]string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler X gewonnen hat.
func PlayerXWins(board [][]string) bool {

	return RowContainsOnly(board, 0, "X") ||
		RowContainsOnly(board, 1, "X") ||
		RowContainsOnly(board, 2, "X") ||
		ColumnContainsOnly(board, 0, "X") ||
		ColumnContainsOnly(board, 1, "X") ||
		ColumnContainsOnly(board, 2, "X") ||
		Diag1ContainsOnly(board, "X") ||
		Diag2ContainsOnly(board, "X")
	/*
	   	for i := range board {
	   		if RowContainsOnly(board, i, "X") || ColumnContainsOnly(board, i, "X") {
	   			return true
	   		}
	   	}

	   return Diag1ContainsOnly(board, "X") || Diag2ContainsOnly(board, "X")
	*/
}

// Erwartet ein Spielfeld. Liefert true, falls Spieler O gewonnen hat.
func PlayerOWins(board [][]string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld. Liefert true, falls das Spiel unentschieden ist.
func Draw(board [][]string) bool {
	// TODO
	return true
}

// Erwartet ein Spielfeld, eine Zeilen- und eine Spaltennummer.
// Liefert true, falls ein Zug an der angegebenen Stelle erlaubt ist.
func MoveAllowed(board [][]string, row, col int) bool {
	// TODO
	return true
}

// Erwartet den String des aktuellen Spielers ("X" oder "O").
// Liefert den String des nächsten Spielers (also den jeweils anderen).
func NextPlayer(currentPlayer string) string {
	// TODO
	return "X"
}
