package main

import (
	"cirks/cirks"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//var players = []string{"red", "blue", "yellow", "green"}
var players = []string{"red"}

var wg sync.WaitGroup

func main() {
	min := 0
	var rollCount int

	var bestRolls []cirks.Roll
	alwaysStart := true

	max := runtime.NumGoroutine() * 4
	c := make(chan []cirks.Roll, max)

	for j := 0; j < 100000; j++ {
		rand.Seed(time.Now().UnixNano())

		wg.Add(max)
		for i := 0; i < max; i++ {
			go cirks.Play(players, true, false, nil, alwaysStart, c, &wg)
		}
		wg.Wait()
		for i := 0; i < max; i++ {
			rolls := <-c

			rollCount = len(rolls)

			if min == 0 {
				min = rollCount
				bestRolls = rolls
				fmt.Println(bestRolls)
			}

			if rollCount <= min {
				min = rollCount
				bestRolls = rolls
				fmt.Println(bestRolls)
			}
		}
	}
	close(c)

	fmt.Println("##########################")

	cirks.Play(players, false, true, bestRolls, alwaysStart, nil, &wg)

	fmt.Println("")
	fmt.Println("STATS")
	fmt.Println("Roll count = " + strconv.Itoa(min))
	fmt.Println("Rolls:")
	fmt.Println(bestRolls)
}
