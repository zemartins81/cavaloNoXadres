package main

import "fmt"

type Position struct {
	x, y int
}

func main() {
	// Recebe o tamanho do tabuleiro
	var boardSize int
	fmt.Print("Entre com o tamanho do tabuleiro: ")
	fmt.Scan(&boardSize)

	// Recebe a posição inicial do cavalo
	var startX, startY int
	fmt.Print("Entre com a posição inicial X do cavalo: ")
	fmt.Scan(&startX)
	fmt.Print("Entre com a posição inicial Y do cavalo: ")
	fmt.Scan(&startY)

	// Lista de movimentos
	var moves []Position

	// Precompute possible next moves for each position
	possibleMoves := make(map[Position][]Position)
	for x := 1; x <= boardSize; x++ {
		for y := 1; y <= boardSize; y++ {
			possibleMoves[Position{x, y}] = calculatePossibleMoves(x, y, boardSize)
		}
	}

	// Loop até chegar no outro lado
	for {
		// Encontra próximo movimento possível
		nextMoves := possibleMoves[Position{startX, startY}]

		// Se não tem mais movimento, acabou
		if len(nextMoves) == 0 {
			break
		}

		// Marca próxima posição como visitada
		nextMove := nextMoves[0]
		moves = append(moves, nextMove)

		if nextMove.y == boardSize {
			break
		}

		// Atualiza posição atual
		startX = nextMove.x
		startY = nextMove.y
	}

	// Exibe movimentos
	fmt.Println(moves)
}

// calculatePossibleMoves calculates all the possible next moves for a given position
func calculatePossibleMoves(x, y, boardSize int) []Position {
	var moves []Position

	possibleOffsets := []Position{
		{1, 2},
		{-1, 2},
		{-2, 1},
		{2, 1},
		{-2, -1},
		{2, -1},
		{-1, -2},
		{1, -2},
	}

	for _, offset := range possibleOffsets {
		nextX := x + offset.x
		nextY := y + offset.y

		if isValid(nextX, nextY, boardSize) {
			moves = append(moves, Position{nextX, nextY})
		}
	}

	return moves
}

// Verifica se a posição é válida
func isValid(x, y, boardSize int) bool {
	return x >= 1 && x <= boardSize && y >= 1 && y <= boardSize
}
