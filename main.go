package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	modEvade = +2
	modDef   = -1
)

func main() {
	flag.Usage = func() {
		fmt.Println(`
sugucalc calculates, if dodging or defending is more optimal for a Suguri-type
100% Orange Juice character (-1 DEF +2 EVD)

usage:
sugucalc ENEMY_ATTACK YOUR_HEALTH
`)
	}
	flag.Parse()
	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	attack := getInt(os.Args[1])
	health := getInt(os.Args[2])
	survivedDefence := 6
	survivedDodge := 6
	for i := 1; i <= 6; i++ {
		if attack >= roll(i, modEvade) && attack >= health {
			survivedDodge--
		}

		dmg := attack - roll(i, modDef)
		if dmg < 1 {
			dmg = 1
		}
		if dmg >= health {
			survivedDefence--
		}
	}

	var re string
	switch {
	case survivedDodge == 0 && survivedDefence == 0:
		re = "you're fucked"
	case survivedDodge >= survivedDefence:
		re = "dodge"
	case survivedDefence > survivedDodge:
		re = "defend"
	}
	fmt.Println(re)
}

func getInt(s string) int {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(i)
}

func roll(i, mod int) int {
	i += mod
	if i < 1 {
		i = 1
	}
	return i
}
