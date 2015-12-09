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
  numbers []int
}

type Dice struct {
  die []int
}

func main() {
  tiles := SetupTiles()
  gameOver := false

  for !gameOver {
    dice := Roll(2)
    PrintBox(tiles, dice)
    moves := AvailableMoves(tiles, dice)
    moves = LegalMoves(tiles, moves)

    if(len(moves) == 0) {
      gameOver = true
      break
    }
    PrintLegalMoves(moves)

    choice := GetChoice()
    CloseTiles(tiles, moves, choice)
  }

  DoGameOver(tiles)

}

func DoGameOver(tiles []*Tile) {
  fmt.Println("")
  fmt.Println("                            ++-- GAME OVER --++")
  fmt.Println("")
}

func GetChoice() int {
  var selection int
  fmt.Printf("Pick move >")
  fmt.Scan(&selection) 
  return selection 
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

func AvailableMoves(tiles []*Tile, dice *Dice) []*Move {
  total := Total(dice)
  moves := Partition(total, nil)
  move := new(Move)
  move.numbers = []int{total}
  moves = append(moves, move)
  return moves
}

func Total(dice *Dice) int {
  total := 0
  for _,num := range dice.die {
    total += num
  }
  return total
}

func Partition(num int, prepend []int) []*Move {
  moves := make([]*Move, 0, 10)
  start := 1

  if prepend != nil {
    start = prepend[len(prepend) - 1] + 1
  }

  for x := start; x <= ((num - 1)/ 2); x++ {
    move := new(Move)

    if prepend != nil {
      move.numbers = append(move.numbers, prepend...)
    } 
    move.numbers = append(move.numbers, x)
    move.numbers = append(move.numbers, num - x)

    moves = append(moves, move)
    moves = append(moves, Partition(num - x, move.numbers[0:(len(move.numbers)-1)])...)
  }

  return moves
}

func LegalMoves(tiles []*Tile, moves []*Move) []*Move {
  outMoves := make([]*Move, 0, 10)
  for _, move := range moves {
    if (LegalMove(tiles, move)) {
      outMoves = append(outMoves, move)
    }
  }
  return outMoves;  
}

func LegalMove(tiles []*Tile, move *Move) bool {
  for _, num := range move.numbers {
    if (num > 9) {
      return false;
    } else if (!tiles[num-1].open) {
      return false;
    }
  }
  return true;
}

func PrintLegalMoves(moves []*Move) {
  for idx, move := range moves {
    PrintMove(idx, move)
  }  
}

func PrintMove(idx int, move *Move) {
  fmt.Printf("%v> ", idx)
  for _, num := range move.numbers {
    fmt.Printf("%v ", num)
  }
  fmt.Printf("\n")
}

func Roll(howMany int) *Dice {
  rand.Seed( time.Now().UTC().UnixNano())
  dice := new(Dice)
  dice.die = make([]int, 0, 5)

  for x:=0; x<howMany; x++ {
    dice.die = append(dice.die, int(rand.Int31n(6) + 1))
  }
  
  return dice
}

func CloseTiles(tiles []*Tile, moves []*Move, choice int) {
  move := moves[choice]

  for _, num := range move.numbers {
    tiles[num-1].open = false;
  }
}

func PrintBox(tiles []*Tile, dice *Dice) {
  PrintTiles(tiles)
  PrintRoll(dice)
}

func PrintRoll(dice *Dice) {
  fmt.Print  ("                   |                                     |\n")
  fmt.Print  ("                   | ")
  fmt.Printf ("       ")
  for _,die := range dice.die {
    fmt.Printf ("[%v] ", die)
  }

  fmt.Print  ("                     |\n")
  fmt.Print  ("                   |                                     |\n")
  fmt.Println("                   +-------------------------------------+")

}

func PrintTiles(tiles []*Tile) {  

  fmt.Println("                   +-------------------------------------+")
  fmt.Print  ("                   | ")

  for i := 0; i < 9; i++ {
    tile := tiles[i]
    if (tile.open) {
      fmt.Printf("[%v] ", tile.number)
    } else {
      fmt.Printf("[-] ")
    }
  }

  fmt.Print  ("|\n")
  fmt.Println("                   +-------------------------------------+")
}