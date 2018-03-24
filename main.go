package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	sum   int64
	pulls int64
	avg   float64
)

func playGame() int {
	var p int
	for p = 1; (rand.Intn(2) != 0) && (p <= 3); p++ {
	}
	return p
}

func pow2(exp int) int64 {
	p2 := int64(1)
	for i := 1; i <= exp; i++ {
		p2 *= 2
	}
	return p2
}

func main() {
	histogram := make(map[int]int)
	for i := 0; i < 1; i++ {
		go func() {
			var max int64
			max = 0
			for {
				pulls++
				tries := playGame()
				var p int64
				if tries == 4 {
					p = 0
				} else {
					p = pow2(tries - 1)
				}
				histogram[tries]++
				sum += p
				avg = float64(sum) / float64(pulls)
				if max < p {
					// fmt.Println("[!] Max", tries, p)
					max = p
				}
				if pulls%10000000 == 0 {
					fmt.Println(pulls, avg, max)
					fmt.Println(histogram)
					fmt.Println("tries =", tries, "won =", p)
					max = 0
				}
			}
		}()
	}

	go func() {
		var previous int64
		previous = 0
		interval := time.Second * 60
		select {
		case <-time.After(interval):
			fmt.Println("Speed", (pulls-previous)/60)
			previous = pulls
		}
	}()

	stopChan := make(chan bool, 1)
	<-stopChan
}
