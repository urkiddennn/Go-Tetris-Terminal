package main

import (
	"fmt"
)

func main() {
	/*	ticker := time.NewTicker(500 * time.Millisecond)
		go func() {
			for t := range ticker.C {
				fmt.Println("Tick at", t)
			}
		}()
		time.Sleep(2 * time.Second)
		ticker.Stop()
	*/
	var width int = 5
	var height int = 5

	for col := range width {
		for row := range height {
			fmt.Println("*")
		}
	}
	fmt.Println("Ticker stop")
}
