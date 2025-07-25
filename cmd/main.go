package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func chessPosToIdx(pos string, board *Board) (int, int, error) {
	if len(pos) < 2 {
		return 0, 0, fmt.Errorf("posição inválida")
	}
	col := int(pos[0] - 'a')
	var row int
	_, err := fmt.Sscanf(pos[1:], "%d", &row)
	if err != nil {
		return 0, 0, fmt.Errorf("linha inválida")
	}
	rowIdx := board.Rows - row
	if col < 0 || col >= board.Columns || rowIdx < 0 || rowIdx >= board.Rows {
		return 0, 0, fmt.Errorf("fora do tabuleiro")
	}
	return rowIdx, col, nil
}

func idxToChessPos(row, col int, board *Board) string {
	return fmt.Sprintf("%c%d", 'A'+col, board.Rows-row)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bem-vindo ao Unvoid Chess Game!")

	var board *Board
	var pieces []*Piece
	var rows, columns int
	for {
		fmt.Print("Digite o número de linhas do tabuleiro (6 a 12): ")
		fmt.Fscanf(reader, "%d\n", &rows)
		fmt.Print("Digite o número de colunas do tabuleiro (6 a 12): ")
		fmt.Fscanf(reader, "%d\n", &columns)
		b, err := NewBoard(rows, columns)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}
		p := InitialPieces(rows, columns)
		board = b
		pieces = p
		fmt.Printf("Tabuleiro criado com %d linhas e %d colunas!\n", board.Rows, board.Columns)
		fmt.Printf("Peças iniciais posicionadas: %d\n", len(pieces))
		PrintBoard(board, pieces)
		break
	}

	turn := White
	var winner *Player
	var selectedPiece *Piece

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nTurno: %s\n", func() string {
			if turn == White {
				return "Branco"
			} else {
				return "Preto"
			}
		}())
		fmt.Print("Digite um comando (help para ver opções): ")
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())

		switch cmd {
		case "help":
			fmt.Println("\nComandos disponíveis:")
			fmt.Println("  help  - mostra esta ajuda")
			fmt.Println("  show  - exibe o tabuleiro atual")
			fmt.Println("  select LINHA COLUNA  - seleciona uma peça na posição informada (ex: select 0 2)")
			fmt.Println("  exit  - encerra o jogo")
		case "show":
			PrintBoard(board, pieces)
		case "status":
			fmt.Printf("Turno atual: %s\n", func() string {
				if turn == White {
					return "Branco"
				} else {
					return "Preto"
				}
			}())
			fmt.Printf("Peças restantes: %d\n", len(pieces))
			if selectedPiece != nil {
				fmt.Printf("Peça selecionada: %s (%d, %d)\n", pieceSymbol(selectedPiece), selectedPiece.Row, selectedPiece.Col)
			} else {
				fmt.Println("Nenhuma peça selecionada.")
			}
		case "reset":
			board, _ = NewBoard(rows, columns)
			pieces = InitialPieces(rows, columns)
			selectedPiece = nil
			turn = White
			winner = nil
			fmt.Println("Jogo reiniciado!")
			PrintBoard(board, pieces)
		case "select":
			fmt.Print("Digite a posição da peça (ex: c1): ")
			scanner.Scan()
			input := strings.Fields(scanner.Text())
			if len(input) != 1 {
				fmt.Println("Por favor, digite apenas uma coordenada (ex: c1).")
				continue
			}
			pos := input[0]
			selRow, selCol, err := chessPosToIdx(pos, board)
			if err != nil {
				fmt.Printf("Posição inválida. O tabuleiro vai de A1 até %c%d.\n", 'A'+board.Columns-1, board.Rows)
				continue
			}
			selectedPiece = nil
			for _, piece := range pieces {
				if piece.Row == selRow && piece.Col == selCol {
					selectedPiece = piece
					break
				}
			}
			if selectedPiece == nil {
				fmt.Println("Não há peça nessa posição.")
				continue
			}
			if selectedPiece.Player != turn {
				cor := "Branco"
				if selectedPiece.Player == Black {
					cor = "Preto"
				}
				fmt.Printf("Você só pode selecionar suas próprias peças. Esta peça pertence ao jogador %s.\n", cor)
				selectedPiece = nil
				continue
			}
			fmt.Printf("Peça selecionada: %s na posição %s\n", pieceSymbol(selectedPiece), pos)
		case "moves":
			if selectedPiece == nil {
				fmt.Println("Nenhuma peça selecionada. Use o comando 'select' primeiro.")
				continue
			}
			moves := PossibleMoves(selectedPiece, board, pieces)
			if len(moves) == 0 {
				fmt.Println("Nenhum movimento possível.")
				continue
			}
			fmt.Printf("Movimentos possíveis para %s na posição %s:\n", pieceSymbol(selectedPiece), idxToChessPos(selectedPiece.Row, selectedPiece.Col, board))
			for _, m := range moves {
				fmt.Printf("-> %s\n", idxToChessPos(m[0], m[1], board))
			}
		case "move":
			if selectedPiece == nil {
				fmt.Println("Nenhuma peça selecionada. Use o comando 'select' primeiro.")
				continue
			}
			fmt.Print("Digite a posição de destino (ex: d3): ")
			scanner.Scan()
			input := strings.Fields(scanner.Text())
			if len(input) != 1 {
				fmt.Println("Por favor, digite apenas uma coordenada (ex: d3).")
				continue
			}
			pos := input[0]
			destRow, destCol, err := chessPosToIdx(pos, board)
			if err != nil {
				fmt.Printf("Posição inválida. O tabuleiro vai de A1 até %c%d.\n", 'A'+board.Columns-1, board.Rows)
				continue
			}
			valid := false
			for _, m := range PossibleMoves(selectedPiece, board, pieces) {
				if m[0] == destRow && m[1] == destCol {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Println("Movimento inválido para esta peça.")
				continue
			}
			for _, p := range pieces {
				if p.Row == destRow && p.Col == destCol && p.Player == turn {
					cor := "Branco"
					if p.Player == Black {
						cor = "Preto"
					}
					fmt.Printf("Você não pode capturar sua própria peça. Esta peça pertence ao jogador %s.\n", cor)
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			capturada := false
			for i, p := range pieces {
				if p.Row == destRow && p.Col == destCol && p.Player != turn {
					if p.Type == ProductOwner {
						winner = &turn
					}
					pieces = append(pieces[:i], pieces[i+1:]...)
					capturada = true
					break
				}
			}
			selectedPiece.Row = destRow
			selectedPiece.Col = destCol
			fmt.Printf("Peça movida para %s%s\n", idxToChessPos(destRow, destCol, board), func() string {
				if capturada {
					return " (captura!)"
				} else {
					return ""
				}
			}())
			PrintBoard(board, pieces)
			if winner != nil {
				fmt.Printf("\nJogo encerrado! Vencedor: %s\n", func() string {
					if *winner == White {
						return "Branco"
					} else {
						return "Preto"
					}
				}())
				return
			}
			if turn == White {
				turn = Black
			} else {
				turn = White
			}
			selectedPiece = nil
			fmt.Printf("Agora é o turno do %s.\n", func() string {
				if turn == White {
					return "Branco"
				} else {
					return "Preto"
				}
			}())
		case "exit":
			fmt.Println("Saindo do jogo. Até logo!")
			return
		default:
			fmt.Println("Comando não reconhecido. Digite 'help' para ver as opções.")
		}
	}
}
