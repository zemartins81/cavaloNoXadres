package main

import "fmt"

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
	var moves []string

	// Loop até chegar no outro lado
	for {
		// Encontra próximo movimento possível
		nextMove := findNextMove(startX, startY, boardSize)

		// Se não tem mais movimento, acabou
		if nextMove == "" {
			break
		}

		// Marca próxima posição como visitada
		x, y := parsePos(nextMove)

		// Adiciona movimento à lista
		moves = append(moves, nextMove)

		if y == boardSize {
			break
		}
		// Atualiza posição atual
		startX = x
		startY = y

	}

	// Exibe movimentos
	fmt.Println(moves)
}

// Encontra o próximo movimento possível
func findNextMove(x, y, boardSize int) string {

	nextY := y + 2
	nextX := x + 1

	if isValid(nextX, nextY, boardSize) {
		return fmt.Sprintf("%d%d", nextX, nextY)
	}


	nextY = y + 2
	nextX = x - 1

	if isValid(nextX, nextY, boardSize) {
		return fmt.Sprintf("%d%d", nextX, nextY)
	}

	nextY = y + 1
	nextX = x - 2

	if isValid(nextX, nextY, boardSize) {
		return fmt.Sprintf("%d%d", nextX, nextY)
	}

	nextY = y - 1
	nextX = x + 2

	if isValid(nextX, nextY, boardSize) {
		return fmt.Sprintf("%d%d", nextX, nextY)
	}

	nextY = y - 1
	nextX = x - 2

	if isValid(nextX, nextY, boardSize) {
		return fmt.Sprintf("%d%d", nextX, nextY)
	}

	return ""
}

// Verifica se a posição é válida
func isValid(x, y, boardSize int) bool {
	if x <= 0 || x > boardSize {
		return false
	}

	if y <= 0 || y > boardSize {
		return false
	}

	return true
}

// Converte posição como "34" para x=3, y=4
func parsePos(pos string) (int, int) {
	x := int(pos[0] - '0')
	y := int(pos[1] - '0')
	return x, y
}
