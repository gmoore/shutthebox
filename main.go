package main


import (
  "fmt"
  "math/rand"
  "time"
)

type Tile struct {
  number int
  open bool
}

type Move struct{
  number1 int
  number2 int
}

func main() {
  tiles := SetupTiles()
  gameOver := false

  for ;!gameOver;{
    PrintTiles(tiles)
    die1 := Roll()
    die2 := Roll()
    PrintRoll(die1, die2)
    moves := AvailableMoves(tiles, die1, die2)

    if(len(moves) == 0) {
      gameOver = true
      break
    }

    choice := GetChoice()
    CloseTiles(tiles, moves, choice)
  }

  DoGameOver(tiles)

}

func DoGameOver(tiles []*Tile) {
  fmt.Println(" ++-- GAME OVER --++")
}

func GetChoice() int {
  var selection int
  fmt.Printf("Pick move >")
  fmt.Scan(&selection) 
  return selection 
}

func PrintRoll(die1 int, die2 int) {
  fmt.Printf(" ROLL-> [%v] [%v]\n", die1, die2)
}

func SetupTiles() []*Tile {
  tiles := make([]*Tile, 0, 10)

  for i := 1; i < 10; i++ {
    tile := new(Tile)
    tile.number = i
    tile.open = true
    tiles = append(tiles, tile)
  }

  return tiles
}

func AvailableMoves(tiles []*Tile, die1 int, die2 int) []*Move {
  total := (die1 + die2)
  moves := make([]*Move, 0, 5)

  for x := 1; x <= ((total - 1)/ 2); x++ {
    if (total-x < 10) {
      move := new(Move)
      move.number1 = x
      move.number2 = total-x
      if(LegalMove(tiles, move)) {
        moves = append(moves, move)
      }
    }
  }

  if (total < 10) {
    move := new(Move)
    move.number1 = total
    if(LegalMove(tiles, move)) {
      moves = append(moves, move)
    }
  }

  fmt.Println("-- MOVES --")
  for idx, move := range moves {
    PrintMove(idx, move)
  }

  return moves
}

func LegalMove(tiles []*Tile, move *Move) bool {
  if (move.number2 > 0) {
    return tiles[move.number1 - 1].open && tiles[move.number2 - 1].open
  } else {
    return tiles[move.number1 - 1].open
  }
}

func PrintMove(idx int, move *Move) {
  if (move.number2 > 0) {
    fmt.Printf("%v> %v %v \n", idx, move.number1, move.number2  )
  } else {
    fmt.Printf("%v> %v \n", idx, move.number1)
  }
}

func Roll() int {
  rand.Seed( time.Now().UTC().UnixNano())
  return int(rand.Int31n(6) + 1)
}

func CloseTiles(tiles []*Tile, moves []*Move, choice int) {
  move := moves[choice]

  for i := 0; i < 9; i++ {
    tile := tiles[i]

    if (tile.number == move.number1 || tile.number == move.number2) {
      tile.open = false
    }
  }
}

func PrintTiles(tiles []*Tile) {  

  fmt.Println(" +-------------------------------------+")
  fmt.Print  (" | ")

  for i := 0; i < 9; i++ {
    tile := tiles[i]

    if (tile.open) {
      fmt.Printf("[%v] ", tile.number)
    } else {
      fmt.Printf("[-] ")
    }
  }

  fmt.Print  ("|\n")
  fmt.Println(" +-------------------------------------+")
  fmt.Println("\n")
}