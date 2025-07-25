package main

import "fmt"

type Board struct {
	Rows    int
	Columns int
}

func NewBoard(rows, columns int) (*Board, error) {
	if rows < 6 || rows > 12 || columns < 6 || columns > 12 {
		return nil, fmt.Errorf("dimensões inválidas: o tabuleiro deve ter entre 6 e 12 linhas e colunas")
	}
	return &Board{
		Rows:    rows,
		Columns: columns,
	}, nil
}

type Player int

const (
	White Player = iota
	Black
)

type PieceType int

const (
	Developer PieceType = iota
	Designer
	ProductOwner
)

type Piece struct {
	Type   PieceType
	Player Player
	Row    int
	Col    int
}

func InitialPieces(rows, columns int) []*Piece {
	pieces := []*Piece{}
	pieces = append(pieces, &Piece{
		Type:   ProductOwner,
		Player: Black,
		Row:    0,
		Col:    columns - 1,
	})
	pieces = append(pieces, &Piece{
		Type:   Developer,
		Player: Black,
		Row:    0,
		Col:    columns - 2,
	})
	pieces = append(pieces, &Piece{
		Type:   Designer,
		Player: Black,
		Row:    0,
		Col:    columns - 3,
	})
	pieces = append(pieces, &Piece{
		Type:   ProductOwner,
		Player: White,
		Row:    rows - 1,
		Col:    0,
	})
	pieces = append(pieces, &Piece{
		Type:   Developer,
		Player: White,
		Row:    rows - 1,
		Col:    1,
	})
	pieces = append(pieces, &Piece{
		Type:   Designer,
		Player: White,
		Row:    rows - 1,
		Col:    2,
	})
	return pieces
}

func PrintBoard(board *Board, pieces []*Piece) {
	grid := make([][]string, board.Rows)
	for i := range grid {
		grid[i] = make([]string, board.Columns)
		for j := range grid[i] {
			grid[i][j] = "□"
		}
	}
	for _, piece := range pieces {
		symbol := pieceSymbol(piece)
		grid[piece.Row][piece.Col] = symbol
	}
	fmt.Print("   ")
	for col := 0; col < board.Columns; col++ {
		fmt.Printf(" %c", 'A'+col)
	}
	fmt.Println()
	for i := 0; i < board.Rows; i++ {
		fmt.Printf("%2d ", board.Rows-i)
		for j := 0; j < board.Columns; j++ {
			fmt.Printf(" %s", grid[i][j])
		}
		fmt.Println()
	}
}

func pieceSymbol(piece *Piece) string {
	switch piece.Type {
	case Developer:
		if piece.Player == White {
			return "♙"
		}
		return "♟"
	case Designer:
		if piece.Player == White {
			return "♘"
		}
		return "♞"
	case ProductOwner:
		if piece.Player == White {
			return "♔"
		}
		return "♚"
	default:
		return "?"
	}
}

func PossibleMoves(piece *Piece, board *Board, pieces []*Piece) [][2]int {
	var moves [][2]int
	isFree := func(row, col int) bool {
		if row < 0 || row >= board.Rows || col < 0 || col >= board.Columns {
			return false
		}
		for _, p := range pieces {
			if p.Row == row && p.Col == col {
				return false
			}
		}
		return true
	}
	isFreeOrEnemy := func(row, col int, player Player) bool {
		if row < 0 || row >= board.Rows || col < 0 || col >= board.Columns {
			return false
		}
		for _, p := range pieces {
			if p.Row == row && p.Col == col {
				return p.Player != player
			}
		}
		return true
	}
	switch piece.Type {
	case Developer:
		dirs := [][2]int{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}
		for _, d := range dirs {
			for dist := 1; dist <= 3; dist++ {
				newRow := piece.Row + d[0]*dist
				newCol := piece.Col + d[1]*dist
				if isFree(newRow, newCol) {
					moves = append(moves, [2]int{newRow, newCol})
				} else {
					break
				}
			}
		}
	case Designer:
		lMoves := [][2]int{
			{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
			{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
		}
		for _, m := range lMoves {
			newRow := piece.Row + m[0]
			newCol := piece.Col + m[1]
			if isFreeOrEnemy(newRow, newCol, piece.Player) {
				moves = append(moves, [2]int{newRow, newCol})
			}
		}
	case ProductOwner:
		dirs := [][2]int{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}
		for _, d := range dirs {
			newRow := piece.Row + d[0]
			newCol := piece.Col + d[1]
			if isFreeOrEnemy(newRow, newCol, piece.Player) {
				moves = append(moves, [2]int{newRow, newCol})
			}
		}
	}
	return moves
}
