package cirks

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var rules = getRules()

type Roll struct {
	dice int
}

func throwDice() int {
	min := 1
	max := 6
	dice := rand.Intn(max-min+1) + min

	return dice
}

func throwAgain(dice int) bool {
	return dice == 1 || dice == 6
}

func Play(players []string, concurrent bool, verbose bool, rerunRolls []Roll, alwaysStart bool, c chan []Roll, wg *sync.WaitGroup) {
	var position = make(map[string]int)
	var rolls []Roll
	var dice int
	turns := 0

	for _, player := range players {
		position[player] = 0
	}

	i := -1
	for {

		for _, player := range players {
			turns++
			for {
				i++

				if len(rerunRolls) > 0 {
					dice = rerunRolls[0].dice
					rerunRolls = rerunRolls[1:]

					if verbose {
						fmt.Println("RERUN THROW  " + strconv.Itoa(dice))
					}
				} else {
					if alwaysStart && i == 0 {
						dice = 6
					} else {
						dice = throwDice()
					}
					if verbose {
						fmt.Println("THROW  " + strconv.Itoa(dice))
					}
				}

				rolls = append(rolls, Roll{dice})

				if position[player] == 0 {
					if verbose {
						fmt.Println("Player " + player + " TRY FOR START throw " + strconv.Itoa(dice))
					}

					if !throwAgain(dice) {
						if verbose {
							fmt.Println("Player " + player + " NO START ALOWED")
						}
						break
					} else {
						position[player] = 1
						if verbose {
							fmt.Println("Player " + player + " SUCCESS START Go to 1")
						}
						continue
					}
				}

				currentPos := position[player]
				if verbose {
					fmt.Println("Player " + player + " current position " + strconv.Itoa(currentPos) + " throw " + strconv.Itoa(dice))
				}
				currentPos = currentPos + dice
				if currentPos > 120 {
					tmp := currentPos - 120
					currentPos = 120 - tmp

					if verbose {
						fmt.Println("Player " + player + " overshoot by " + strconv.Itoa(tmp) + " current position " + strconv.Itoa(currentPos))
					}
				}

				if verbose {
					fmt.Println("Player " + player + " new position " + strconv.Itoa(currentPos))
				}

				rule, ok := rules[currentPos]
				if ok {
					currentPos = getTo(rule)
					if verbose {
						fmt.Println("Player " + player + " Stepped on Rule. (" + getRuleName(rule) + ")")
					}
				}

				position[player] = currentPos
				if currentPos == 120 {
					if verbose {
						fmt.Println("Player " + player + " WON")
					}

					if concurrent {
						c <- rolls
						wg.Done()
					}

					return
				}

				if !throwAgain(dice) {
					break
				}

				if verbose {
					fmt.Println("Player " + player + " rethrow again (because " + strconv.Itoa(dice) + ")")
				}
			}
		}
	}
}
