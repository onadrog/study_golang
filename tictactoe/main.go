package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grid [3][3]string

func yaWin(player string) {
	msg := fmt.Sprintf("Oi! %s Ya win son !", player)
	fmt.Println(msg)
	os.Exit(0)
}

func (g *Grid) areYaWinningSon(line, column int, player string) error{
	max := len(g)
  if g[line][column] == "X" ||g[line][column] == "O"{
    return fmt.Errorf("This numbah is already taken")
  }
	g[line][column] = player

	// check diags
	if line == column {
		for i := 0; i < max; i++ {
			if g[i][i] != player {
				break
			}
			if i == max-1 {
				yaWin(player)
			}
		}
	}

	// check anti diags

	if line+column == max-1 {
		for i := 0; i < max; i++ {
			if g[i][(max-1)-i] != player {
				break
			}
			if i == max-1 {
				yaWin(player)
			}
		}
	}

	// check lines
	for i := 0; i < max; i++ {
		if g[line][i] != player {
			break
		}
		if i == max-1 {
			yaWin(player)
		}
	}

	// check columns
	for i := 0; i < max; i++ {
		if g[i][column] != player {
			break
		}
		if i == max-1 {
			yaWin(player)
		}
	}
  return nil
}

func (g *Grid) promptGameRecursive() {
	res := ""
	player := "X"
	count := 0
	for {
    if count == len(g)*len(g) {
			fmt.Println("that a draw !")
			os.Exit(2)
		}
		count++
		for i := 0; i < len(g); i++ {
			fmt.Println(g[i])
		}
		fmt.Println(player, "Choose a numbah in da grid ")

		_, err := fmt.Scanln(&res)
		if err != nil {
			log.Fatal(err)
		}

		value, err := strconv.Atoi(res)
		if err != nil {
			fmt.Println("BRUH! Enter a numbah !")
			continue
		}

		if value < 0 || value > 8 {
			fmt.Println("BRUH! Did you really see your numbah in da grid !")
			continue
		}

		line := value / 3
		column := value % 3

    if err := g.areYaWinningSon(line, column, player); err != nil {
      fmt.Println(err)
      count --
      continue
    }
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
		
	}
}

func (g *Grid) fillGrid() {
	count := 0
	for r, c := range g {
		for j := range c {
			g[r][j] = strconv.Itoa(count)
			count++
		}
	}
}

func main() {
	g := &Grid{}
	g.fillGrid()
	g.promptGameRecursive()
}
