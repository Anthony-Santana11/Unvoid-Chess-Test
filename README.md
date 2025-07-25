Claro! Aqui está o README traduzido para o português logo após a versão em inglês:

---

# Unvoid Chess Game (Candidate Test)

A command-line chess-inspired game implemented in Go, created as a technical test for Unvoid.

## Features

- Customizable board size (6x6 to 12x12)
- Three piece types: Developer, Designer, Product Owner
- Turn-based gameplay (White vs Black)
- Each piece has unique movement and capture rules
- CLI interface with commands for selecting, moving, and viewing the board
- Clear error messages and robust input validation

## How to Run

### With Go installed

```sh
cd cmd
go run .
```

### With Docker

```sh
docker build -t candidate-test .
docker run --rm -it candidate-test
```

## How to Play

- The game starts by asking for the board size.
- Use the CLI commands to interact:

| Command         | Description                                      |
|-----------------|--------------------------------------------------|
| help            | Show available commands                          |
| show            | Display the current board                        |
| select <pos>    | Select a piece at position (e.g., select c1)     |
| moves           | Show possible moves for the selected piece       |
| move <pos>      | Move the selected piece to position (e.g., d3)   |
| status          | Show current turn and selected piece             |
| reset           | Restart the game                                 |
| exit            | Exit the game                                    |

- The board uses chess notation: columns (A, B, C, ...) and rows (1, 2, 3, ...).

## Piece Rules

- **Developer (♙/♟):** Moves up to 3 squares in any direction, but only captures by jumping over an opponent's piece (like in checkers).
- **Designer (♘/♞):** Moves in an "L" shape (like a knight in chess), captures by landing on an opponent's piece.
- **Product Owner (♔/♚):** Moves 1 square in any direction, captures by landing on an opponent's piece.

## Objective

Capture the opponent's Product Owner to win the game!

## Example

```
   A B C D E F
 8 □ □ ♘ □ □ □
 7 □ □ □ □ ♚ □
 6 □ □ □ ♙ □ □
 5 □ ♞ □ □ □ □
 4 □ □ □ □ □ □
 3 □ □ □ □ □ □
 2 □ □ □ □ □ □
 1 ♔ □ □ □ □ □
Turn: White
```

## Notes

- Only your own pieces can be selected and moved.
- Invalid moves and selections are clearly explained.
- The game is fully playable via the terminal.

---

# Unvoid Chess Game (Teste de Candidato)

Um jogo de tabuleiro inspirado no xadrez, rodando no terminal, implementado em Go como teste técnico para a Unvoid.

## Funcionalidades

- Tamanho do tabuleiro customizável (6x6 até 12x12)
- Três tipos de peças: Developer, Designer, Product Owner
- Jogo por turnos (Branco vs Preto)
- Cada peça tem regras únicas de movimento e captura
- Interface CLI com comandos para selecionar, mover e visualizar o tabuleiro
- Mensagens de erro claras e validação robusta de entrada

## Como Rodar

### Com Go instalado

```sh
cd cmd
go run .
```

### Com Docker

```sh
docker build -t candidate-test .
docker run --rm -it candidate-test
```

## Como Jogar

- O jogo começa pedindo o tamanho do tabuleiro.
- Use os comandos no terminal para interagir:

| Comando         | Descrição                                         |
|-----------------|---------------------------------------------------|
| help            | Mostra os comandos disponíveis                    |
| show            | Exibe o tabuleiro atual                           |
| select <pos>    | Seleciona uma peça na posição (ex: select c1)     |
| moves           | Mostra os movimentos possíveis da peça selecionada|
| move <pos>      | Move a peça selecionada para a posição (ex: d3)   |
| status          | Mostra o turno atual e a peça selecionada         |
| reset           | Reinicia o jogo                                   |
| exit            | Encerra o jogo                                    |

- O tabuleiro usa notação de xadrez: colunas (A, B, C, ...) e linhas (1, 2, 3, ...).

## Regras das Peças

- **Developer (♙/♟):** Move até 3 casas em qualquer direção, mas só captura pulando sobre uma peça adversária (como na dama).
- **Designer (♘/♞):** Move em “L” (como o cavalo do xadrez), captura ao cair na casa da peça adversária.
- **Product Owner (♔/♚):** Move 1 casa em qualquer direção, captura ao cair na casa da peça adversária.

## Objetivo

Capture o Product Owner do adversário para vencer o jogo!

## Exemplo

```
   A B C D E F
 8 □ □ ♘ □ □ □
 7 □ □ □ □ ♚ □
 6 □ □ □ ♙ □ □
 5 □ ♞ □ □ □ □
 4 □ □ □ □ □ □
 3 □ □ □ □ □ □
 2 □ □ □ □ □ □
 1 ♔ □ □ □ □ □
Turno: Branco
```

## Observações

- Só é possível selecionar e mover suas próprias peças.
- Movimentos e seleções inválidas são explicados claramente.
- O jogo é totalmente jogável via terminal.
