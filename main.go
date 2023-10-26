package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const size = 10

var arr = [size][size]byte{}
var arr2 = [size][size]byte{}

func main() {
	setupArr()

	s := 0
	for s < 10 {
		printArr()
		step()
		s++
	}
}

func printArr() {
	for i := 0; i < size; i++ {
		var line string
		for j := 0; j < size; j++ {
			if arr[i][j] == 255 {
				line = line + "#"
			} else {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func setupArr() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr[i][j] = byte((rand.Int() % 2) * 255)
		}
	}
}

func step() {
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			wg.Add(1)
			go calculatePixel(i, j, &wg)
		}
	}
	wg.Wait()
	arr = arr2
}

func calculatePixel(i int, j int, wg *sync.WaitGroup) {
	sum := 0
	defer wg.Done()
	for x := max(i-1, 0); x < min(size-1, i+2); x++ {
		for y := max(j-1, 0); y < min(size-1, j+2); y++ {
			if x == i && y == j {
				continue
			}
			if arr[x][y] == 255 {
				sum++
			}
		}
	}
	if sum < 2 || sum > 3 {
		arr2[i][j] = 0
	} else if sum == 3 {
		arr2[i][j] = 255
	} else {
		arr2[i][j] = arr[i][j]
	}
}
